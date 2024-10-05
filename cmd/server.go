package cmd

import (
	"context"
	"fmt"

	"github.com/jbub/foxesscloud"
	"github.com/jbub/foxesscloud_exporter/internal/collector"
	"github.com/jbub/foxesscloud_exporter/internal/config"
	"github.com/jbub/foxesscloud_exporter/internal/server"

	"github.com/oklog/run"
	"github.com/prometheus/common/version"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var Server = &cli.Command{
	Name:   "server",
	Usage:  "Starts exporter server.",
	Action: runServer,
}

func runServer(ctx *cli.Context) error {
	cfg := config.LoadFromCLI(ctx)
	log, err := newLogger(cfg.LogLevel)
	if err != nil {
		return fmt.Errorf("could not create logger: %v", err)
	}

	client, err := foxesscloud.NewClient(foxesscloud.Config{
		Token:     cfg.APIToken,
		UserAgent: collector.Name,
	})
	if err != nil {
		return fmt.Errorf("could not create client: %v", err)
	}

	var g run.Group

	exp, err := collector.New(cfg, log, client)
	if err != nil {
		return fmt.Errorf("could not create exporter: %v", err)
	}

	g.Add(func() error {
		return exp.Start()
	}, func(err error) {
		exp.Shutdown()
	})

	srv := server.New(cfg, exp)
	g.Add(func() error {
		return srv.Run()
	}, func(err error) {
		_ = srv.Shutdown(context.Background())
	})

	log.Info("Starting exporter",
		zap.String("listen_addr", cfg.ListenAddress),
		zap.String("telemetry_path", cfg.TelemetryPath),
		zap.String("build_context", version.BuildContext()),
	)

	return g.Run()
}

func newLogger(level string) (*zap.Logger, error) {
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	logCfg := zap.NewProductionConfig()
	logCfg.Level = lvl
	return logCfg.Build()
}
