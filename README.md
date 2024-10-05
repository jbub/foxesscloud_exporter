# Fox ESS Cloud exporter
[![Build Status](https://github.com/jbub/foxesscloud_exporter/actions/workflows/go.yml/badge.svg)][build]
[![Docker Pulls](https://img.shields.io/docker/pulls/jbub/foxesscloud_exporter.svg?maxAge=604800)][hub]
[![Go Report Card](https://goreportcard.com/badge/github.com/jbub/foxesscloud_exporter)][goreportcard]

Prometheus exporter for Fox ESS Cloud inverter metrics.

## Docker

Metrics are by default exposed on http server running on port `9561` under the `/metrics` path.

```bash
docker run \ 
  --detach \ 
  --env "INVERTERS=my-inverter-sn-1,my-inverter-sn-2" \
  --env "API_TOKEN=my-foxess-api-token" \
  --publish "9561:9561" \
  --name "foxesscloud_exporter" \
  jbub/foxesscloud_exporter
```

## Default constant prometheus labels

In order to provide default prometheus constant labels you can use the `DEFAULT_LABELS` environment variable.
Labels can be set in this format `instance=pg1 env=dev`. Provided labels will be added to all the metrics.

[build]: https://github.com/jbub/foxesscloud_exporter/actions/workflows/go.yml
[hub]: https://hub.docker.com/r/jbub/foxesscloud_exporter
[goreportcard]: https://goreportcard.com/report/github.com/jbub/foxesscloud_exporter
