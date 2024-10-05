package collector

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func buildMetrics() []metric {
	return []metric{
		{
			name:    "ambient_temperature_celsius",
			help:    "Internal temperature of the inverter in celsius.",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.AmbientTemperature },
		},
		{
			name:    "boost_temperature_celsius",
			help:    "Boost temperature of the inverter in celsius.",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.BoostTemperature },
		},
		{
			name:    "inverter_temperature_celsius",
			help:    "Temperature of the inverter in celsius.",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.InverterTemperature },
		},
		{
			name:    "generated_power_today_kwh",
			help:    "Today generated power",
			valType: prometheus.CounterValue,
			eval:    func(data metricData) float64 { return data.TodayGeneratedPower },
		},
		{
			name:    "generated_power_total_kwh",
			help:    "Total generated power",
			valType: prometheus.CounterValue,
			eval:    func(data metricData) float64 { return data.TotalGeneratedPower },
		},
		{
			name:    "photovoltaic_power_kwh",
			help:    "Photovoltaic power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PhotovoltaicPower },
		},
		{
			name:    "load_power_kw",
			help:    "Load power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.LoadPower },
		},
		{
			name:    "output_power_kw",
			help:    "Output power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.OutputPower },
		},
		{
			name:    "grid_consumption_power_kw",
			help:    "Grid consumption power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.GridConsumptionPower },
		},
		{
			name:    "pv1_voltage_v",
			help:    "PV1 voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV1Voltage },
		},
		{
			name:    "pv1_current_amp",
			help:    "PV1 current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV1Current },
		},
		{
			name:    "pv1_power_kw",
			help:    "PV1 power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV1Power },
		},
		{
			name:    "pv2_voltage_v",
			help:    "PV2 voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV2Voltage },
		},
		{
			name:    "pv2_current_amp",
			help:    "PV2 current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV2Current },
		},
		{
			name:    "pv2_power_kw",
			help:    "PV2 power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV2Power },
		},
		{
			name:    "pv3_voltage_v",
			help:    "PV3 voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV3Voltage },
		},
		{
			name:    "pv3_current_amp",
			help:    "PV3 current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV3Current },
		},
		{
			name:    "pv3_power_kw",
			help:    "PV3 power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV3Power },
		},
		{
			name:    "pv4_voltage_v",
			help:    "PV4 voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV4Voltage },
		},
		{
			name:    "pv4_current_amp",
			help:    "PV4 current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV4Current },
		},
		{
			name:    "pv4_power_kw",
			help:    "PV4 power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.PV4Power },
		},
		{
			name:    "reference_frequency_hz",
			help:    "Reference frequency",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.ReferenceFrequency },
		},
		{
			name:    "reference_voltage_v",
			help:    "Reference voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.ReferenceVoltage },
		},
		{
			name:    "reference_current_amp",
			help:    "Reference current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.ReferenceCurrent },
		},
		{
			name:    "reference_power_kw",
			help:    "Reference power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.ReferencePower },
		},
		{
			name:    "secondary_frequency_hz",
			help:    "Secondary frequency",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.SecondaryFrequency },
		},
		{
			name:    "secondary_voltage_v",
			help:    "Secondary voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.SecondaryVoltage },
		},
		{
			name:    "secondary_current_amp",
			help:    "Secondary current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.SecondaryCurrent },
		},
		{
			name:    "secondary_power_kw",
			help:    "Secondary power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.SecondaryPower },
		},
		{
			name:    "tertiary_frequency_hz",
			help:    "Tertiary frequency",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.TertiaryFrequency },
		},
		{
			name:    "tertiary_voltage_v",
			help:    "Tertiary voltage",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.TertiaryVoltage },
		},
		{
			name:    "tertiary_current_amp",
			help:    "Tertiary current",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.TertiaryCurrent },
		},
		{
			name:    "tertiary_power_kw",
			help:    "Tertiary power",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.TertiaryPower },
		},
		{
			name:    "fault_count",
			help:    "Number of errors reported.",
			valType: prometheus.CounterValue,
			eval:    func(data metricData) float64 { return data.FaultCount },
		},
		{
			name:    "running_state",
			help:    "Running state.",
			valType: prometheus.GaugeValue,
			eval:    func(data metricData) float64 { return data.RunningState },
		},
		{
			name:    "last_updated_timestamp_seconds",
			help:    "Timestamp of the last update in seconds.",
			valType: prometheus.CounterValue,
			eval:    func(data metricData) float64 { return float64(data.UpdateTime.Unix()) },
		},
	}
}

type metricData struct {
	InverterSN   string
	RunningState float64
	FaultCount   float64

	AmbientTemperature  float64
	BoostTemperature    float64
	InverterTemperature float64

	PhotovoltaicPower    float64
	FeedInPower          float64
	TodayGeneratedPower  float64
	TotalGeneratedPower  float64
	LoadPower            float64
	OutputPower          float64
	GridConsumptionPower float64

	PV1Power   float64
	PV1Voltage float64
	PV1Current float64

	PV2Power   float64
	PV2Voltage float64
	PV2Current float64

	PV3Power   float64
	PV3Voltage float64
	PV3Current float64

	PV4Power   float64
	PV4Voltage float64
	PV4Current float64

	ReferencePower     float64
	ReferenceVoltage   float64
	ReferenceCurrent   float64
	ReferenceFrequency float64

	SecondaryPower     float64
	SecondaryVoltage   float64
	SecondaryCurrent   float64
	SecondaryFrequency float64

	TertiaryPower     float64
	TertiaryVoltage   float64
	TertiaryCurrent   float64
	TertiaryFrequency float64

	UpdateTime time.Time
}
