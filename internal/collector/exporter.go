package collector

import (
	"context"
	"fmt"
	"maps"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jbub/foxesscloud"
	"github.com/jbub/foxesscloud_exporter/internal/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/collectors/version"
	"go.uber.org/zap"
)

const (
	Name = "foxesscloud_exporter"
)

var (
	_ prometheus.Collector = &Exporter{}
)

type metric struct {
	name    string
	help    string
	valType prometheus.ValueType
	eval    func(data metricData) float64
}

func (m metric) desc(constLabels prometheus.Labels) *prometheus.Desc {
	return prometheus.NewDesc(prometheus.BuildFQName("foxesscloud", "", m.name), m.help, nil, constLabels)
}

type Exporter struct {
	log         *zap.Logger
	inverters   []string
	constLabels prometheus.Labels
	metrics     []metric
	client      *foxesscloud.Client
	data        atomic.Pointer[[]metricData]
	interval    time.Duration
	timeout     time.Duration
	done        chan struct{}
}

func New(cfg config.Config, log *zap.Logger, client *foxesscloud.Client) (*Exporter, error) {
	if len(cfg.Inverters) == 0 {
		return nil, fmt.Errorf("no inverters defined")
	}

	return &Exporter{
		log:         log,
		inverters:   cfg.Inverters,
		constLabels: parseLabels(cfg.DefaultLabels),
		metrics:     buildMetrics(),
		client:      client,
		interval:    cfg.APIFetchInterval,
		timeout:     cfg.APIFetchTimeout,
		done:        make(chan struct{}, 1),
	}, nil
}

func (e *Exporter) Start() error {
	ctx := context.Background()
	data, err := e.fetchInverters(ctx)
	if err != nil {
		return fmt.Errorf("could not fetch inverter data: %w", err)
	}
	e.data.Store(&data)

	tick := time.NewTicker(e.interval)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			data, err := e.fetchInverters(ctx)
			if err != nil {
				e.log.Error("could not fetch inverter data", zap.Error(err))
				continue
			}
			e.data.Store(&data)
		case <-e.done:
			return nil
		}
	}
}

func (e *Exporter) Shutdown() {
	close(e.done)
}

func (e *Exporter) Describe(descs chan<- *prometheus.Desc) {
	for _, inverterSN := range e.inverters {
		labels := e.buildLabels(inverterSN)
		for _, m := range e.metrics {
			descs <- m.desc(labels)
		}
	}
}

func (e *Exporter) Collect(metrics chan<- prometheus.Metric) {
	data := e.data.Load()
	if data == nil {
		return
	}

	for _, m := range e.metrics {
		for _, d := range *data {
			metrics <- prometheus.MustNewConstMetric(
				m.desc(e.buildLabels(d.InverterSN)),
				m.valType,
				m.eval(d),
			)
		}
	}
}

const (
	inverterSNLabel = "inverter_sn"
)

func (e *Exporter) buildLabels(inverterSN string) prometheus.Labels {
	if len(e.constLabels) == 0 {
		return prometheus.Labels{inverterSNLabel: inverterSN}
	}
	labels := maps.Clone(e.constLabels)
	labels[inverterSNLabel] = inverterSN
	return labels
}

func (e *Exporter) fetchInverters(ctx context.Context) ([]metricData, error) {
	ctx, cancel := context.WithTimeout(ctx, e.timeout)
	defer cancel()

	res := make([]metricData, 0, len(e.inverters))
	for _, inverterSN := range e.inverters {
		data, err := e.fetchInverterData(ctx, inverterSN)
		if err != nil {
			return nil, err
		}
		res = append(res, data)
	}
	return res, nil
}

func (e *Exporter) fetchInverterData(ctx context.Context, inverterSN string) (metricData, error) {
	e.log.Debug("fetching inverter data", zap.String("inverter_sn", inverterSN))

	data, err := e.client.Inverters.GetRealtimeData(ctx, foxesscloud.GetInverterRealtimeDataOptions{
		InverterSN: inverterSN,
	})
	if err != nil {
		return metricData{}, err
	}

	e.log.Debug("fetched inverter data", zap.String("inverter_sn", inverterSN), zap.Int("num_items", len(data.Items)))

	if len(data.Items) == 0 {
		return metricData{}, fmt.Errorf("no data")
	}
	item := data.Items[0]

	d := metricData{
		InverterSN: inverterSN,
		UpdateTime: item.Time.Time,
	}

	for _, dataItem := range item.Datas {
		switch dataItem.Variable {
		case foxesscloud.VariableGeneration:
			d.TotalGeneratedPower = dataItem.Value.Value
		case foxesscloud.VariableTodayYield:
			d.TodayGeneratedPower = dataItem.Value.Value
		case foxesscloud.VariableFeedinPower:
			d.FeedInPower = dataItem.Value.Value
		case foxesscloud.VariablePvPower:
			d.PhotovoltaicPower = dataItem.Value.Value
		case foxesscloud.VariableLoadsPower:
			d.LoadPower = dataItem.Value.Value
		case foxesscloud.VariableGenerationPower:
			d.OutputPower = dataItem.Value.Value
		case foxesscloud.VariableGridConsumptionPower:
			d.GridConsumptionPower = dataItem.Value.Value
		case foxesscloud.VariableAmbientTemperation:
			d.AmbientTemperature = dataItem.Value.Value
		case foxesscloud.VariableBoostTemperation:
			d.BoostTemperature = dataItem.Value.Value
		case foxesscloud.VariableInvTemperation:
			d.InverterTemperature = dataItem.Value.Value
		case foxesscloud.VariablePv1Volt:
			d.PV1Voltage = dataItem.Value.Value
		case foxesscloud.VariablePv1Current:
			d.PV1Current = dataItem.Value.Value
		case foxesscloud.VariablePv1Power:
			d.PV1Power = dataItem.Value.Value
		case foxesscloud.VariablePv2Volt:
			d.PV2Voltage = dataItem.Value.Value
		case foxesscloud.VariablePv2Current:
			d.PV2Current = dataItem.Value.Value
		case foxesscloud.VariablePv2Power:
			d.PV2Power = dataItem.Value.Value
		case foxesscloud.VariablePv3Volt:
			d.PV3Voltage = dataItem.Value.Value
		case foxesscloud.VariablePv3Current:
			d.PV3Current = dataItem.Value.Value
		case foxesscloud.VariablePv3Power:
			d.PV3Power = dataItem.Value.Value
		case foxesscloud.VariablePv4Volt:
			d.PV4Voltage = dataItem.Value.Value
		case foxesscloud.VariablePv4Current:
			d.PV4Current = dataItem.Value.Value
		case foxesscloud.VariablePv4Power:
			d.PV4Power = dataItem.Value.Value
		case foxesscloud.VariableRPower:
			d.ReferencePower = dataItem.Value.Value
		case foxesscloud.VariableRVolt:
			d.ReferenceVoltage = dataItem.Value.Value
		case foxesscloud.VariableRCurrent:
			d.ReferenceCurrent = dataItem.Value.Value
		case foxesscloud.VariableRFreq:
			d.ReferenceFrequency = dataItem.Value.Value
		case foxesscloud.VariableSPower:
			d.SecondaryPower = dataItem.Value.Value
		case foxesscloud.VariableSVolt:
			d.SecondaryVoltage = dataItem.Value.Value
		case foxesscloud.VariableSCurrent:
			d.SecondaryCurrent = dataItem.Value.Value
		case foxesscloud.VariableSFreq:
			d.SecondaryFrequency = dataItem.Value.Value
		case foxesscloud.VariableTPower:
			d.TertiaryPower = dataItem.Value.Value
		case foxesscloud.VariableTVolt:
			d.TertiaryVoltage = dataItem.Value.Value
		case foxesscloud.VariableTCurrent:
			d.TertiaryCurrent = dataItem.Value.Value
		case foxesscloud.VariableTFreq:
			d.TertiaryFrequency = dataItem.Value.Value
		case foxesscloud.VariableRunningState:
			d.RunningState = dataItem.Value.Value
		case foxesscloud.VariableCurrentFaultCount:
			d.FaultCount = dataItem.Value.Value
		}
	}
	return d, nil
}

func parseLabels(s string) prometheus.Labels {
	if s == "" {
		return nil
	}

	items := strings.Split(s, " ")
	res := make(prometheus.Labels, len(items))
	for _, item := range items {
		if item == "" {
			continue
		}
		if parts := strings.SplitN(item, "=", 2); len(parts) == 2 {
			res[parts[0]] = parts[1]
		}
	}
	return res
}

func NewRegistry(exp *Exporter) prometheus.Gatherer {
	reg := prometheus.NewRegistry()
	reg.MustRegister(version.NewCollector(Name))
	reg.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{
		Namespace:    "",
		ReportErrors: false,
	}))
	reg.MustRegister(collectors.NewGoCollector())
	reg.MustRegister(exp)
	return reg
}
