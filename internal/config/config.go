package config

import (
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

func LoadFromCLI(ctx *cli.Context) Config {
	return Config{
		LogLevel:         ctx.String("log-level"),
		ListenAddress:    ctx.String("web.listen-address"),
		TelemetryPath:    ctx.String("web.telemetry-path"),
		Inverters:        parseInverters(ctx.String("inverters")),
		APIToken:         ctx.String("api-token"),
		APIFetchInterval: ctx.Duration("api-fetch-interval"),
		APIFetchTimeout:  ctx.Duration("api-fetch-timeout"),
		DefaultLabels:    ctx.String("default-labels"),
	}
}

type Config struct {
	LogLevel         string
	ListenAddress    string
	TelemetryPath    string
	Inverters        []string
	APIToken         string
	APIFetchInterval time.Duration
	APIFetchTimeout  time.Duration
	DefaultLabels    string
}

func parseInverters(inverters string) []string {
	split := strings.Split(inverters, ",")
	res := make([]string, 0, len(split))
	for _, s := range split {
		if clean := strings.TrimSpace(s); clean != "" {
			res = append(res, clean)
		}
	}
	return res
}
