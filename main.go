package main

import (
	"log"
	"os"
	"time"

	"github.com/jbub/foxesscloud_exporter/cmd"

	"github.com/prometheus/common/version"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "foxesscloud_exporter",
		Usage: "foxesscloud_exporter",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "web.listen-address",
				Usage:   "Address on which to expose metrics and web interface.",
				EnvVars: []string{"WEB_LISTEN_ADDRESS"},
				Value:   ":9561",
			},
			&cli.StringFlag{
				Name:    "web.telemetry-path",
				Usage:   "Path under which to expose metrics.",
				EnvVars: []string{"WEB_TELEMETRY_PATH"},
				Value:   "/metrics",
			},
			&cli.StringFlag{
				Name:     "inverters",
				Usage:    "Comma separated list of inverter serial numbers.",
				EnvVars:  []string{"INVERTERS"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "api-token",
				Usage:    "API token for the Fox ESS API.",
				EnvVars:  []string{"API_TOKEN"},
				Required: true,
			},
			&cli.DurationFlag{
				Name:    "api-fetch-interval",
				Usage:   "How often to fetch the API.",
				EnvVars: []string{"API_FETCH_INTERVAL"},
				Value:   time.Second * 10,
			},
			&cli.DurationFlag{
				Name:    "api-fetch-timeout",
				Usage:   "How long to wait for API fetch response.",
				EnvVars: []string{"API_FETCH_TIMEOUT"},
				Value:   time.Second * 5,
			},
			&cli.StringFlag{
				Name:    "log-level",
				Usage:   "Default log level.",
				EnvVars: []string{"LOG_LEVEL"},
				Value:   "info",
			},
			&cli.StringFlag{
				Name:    "default-labels",
				Usage:   "Default prometheus labels applied to all metrics. Format: label1=value1 label2=value2",
				EnvVars: []string{"DEFAULT_LABELS"},
			},
		},
		Commands: []*cli.Command{
			cmd.Server,
		},
		Version: version.Info(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
