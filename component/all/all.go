// Package all imports all known component packages.
package all

import (
	_ "github.com/grafana/agent/component/discovery/aws"                            // Import discovery.aws.ec2 and discovery.aws.lightsail
	_ "github.com/grafana/agent/component/discovery/azure"                          // Import discovery.azure
	_ "github.com/grafana/agent/component/discovery/consul"                         // Import discovery.consul
	_ "github.com/grafana/agent/component/discovery/digitalocean"                   // Import discovery.digitalocean
	_ "github.com/grafana/agent/component/discovery/dns"                            // Import discovery.dns
	_ "github.com/grafana/agent/component/discovery/docker"                         // Import discovery.docker
	_ "github.com/grafana/agent/component/discovery/eureka"                         // Import discovery.eureka
	_ "github.com/grafana/agent/component/discovery/file"                           // Import discovery.file
	_ "github.com/grafana/agent/component/discovery/gce"                            // Import discovery.gce
	_ "github.com/grafana/agent/component/discovery/http"                           // Import discovery.http
	_ "github.com/grafana/agent/component/discovery/kubelet"                        // Import discovery.kubelet
	_ "github.com/grafana/agent/component/discovery/kubernetes"                     // Import discovery.kubernetes
	_ "github.com/grafana/agent/component/discovery/nomad"                          // Import discovery.nomad
	_ "github.com/grafana/agent/component/discovery/openstack"                      // Import discovery.openstack
	_ "github.com/grafana/agent/component/discovery/relabel"                        // Import discovery.relabel
	_ "github.com/grafana/agent/component/discovery/uyuni"                          // Import discovery.uyuni
	_ "github.com/grafana/agent/component/local/file"                               // Import local.file
	_ "github.com/grafana/agent/component/local/file_match"                         // Import local.file_match
	_ "github.com/grafana/agent/component/loki/echo"                                // Import loki.echo
	_ "github.com/grafana/agent/component/loki/process"                             // Import loki.process
	_ "github.com/grafana/agent/component/loki/relabel"                             // Import loki.relabel
	_ "github.com/grafana/agent/component/loki/source/api"                          // Import loki.source.api
	_ "github.com/grafana/agent/component/loki/source/aws_firehose"                 // Import loki.source.awsfirehose
	_ "github.com/grafana/agent/component/loki/source/azure_event_hubs"             // Import loki.source.azure_event_hubs
	_ "github.com/grafana/agent/component/loki/source/cloudflare"                   // Import loki.source.cloudflare
	_ "github.com/grafana/agent/component/loki/source/docker"                       // Import loki.source.docker
	_ "github.com/grafana/agent/component/loki/source/file"                         // Import loki.source.file
	_ "github.com/grafana/agent/component/loki/source/gcplog"                       // Import loki.source.gcplog
	_ "github.com/grafana/agent/component/loki/source/gelf"                         // Import loki.source.gelf
	_ "github.com/grafana/agent/component/loki/source/heroku"                       // Import loki.source.heroku
	_ "github.com/grafana/agent/component/loki/source/journal"                      // Import loki.source.journal
	_ "github.com/grafana/agent/component/loki/source/kafka"                        // Import loki.source.kafka
	_ "github.com/grafana/agent/component/loki/source/kubernetes"                   // Import loki.source.kubernetes
	_ "github.com/grafana/agent/component/loki/source/kubernetes_events"            // Import loki.source.kubernetes_events
	_ "github.com/grafana/agent/component/loki/source/podlogs"                      // Import loki.source.podlogs
	_ "github.com/grafana/agent/component/loki/source/syslog"                       // Import loki.source.syslog
	_ "github.com/grafana/agent/component/loki/source/windowsevent"                 // Import loki.source.windowsevent
	_ "github.com/grafana/agent/component/loki/write"                               // Import loki.write
	_ "github.com/grafana/agent/component/mimir/rules/kubernetes"                   // Import mimir.rules.kubernetes
	_ "github.com/grafana/agent/component/module/file"                              // Import module.file
	_ "github.com/grafana/agent/component/module/git"                               // Import module.git
	_ "github.com/grafana/agent/component/module/http"                              // Import module.http
	_ "github.com/grafana/agent/component/module/string"                            // Import module.string
	_ "github.com/grafana/agent/component/otelcol/auth/basic"                       // Import otelcol.auth.basic
	_ "github.com/grafana/agent/component/otelcol/auth/bearer"                      // Import otelcol.auth.bearer
	_ "github.com/grafana/agent/component/otelcol/auth/headers"                     // Import otelcol.auth.headers
	_ "github.com/grafana/agent/component/otelcol/auth/oauth2"                      // Import otelcol.auth.oauth2
	_ "github.com/grafana/agent/component/otelcol/auth/sigv4"                       // Import otelcol.auth.sigv4
	_ "github.com/grafana/agent/component/otelcol/exporter/jaeger"                  // Import otelcol.exporter.jaeger
	_ "github.com/grafana/agent/component/otelcol/exporter/loadbalancing"           // Import otelcol.exporter.loadbalancing
	_ "github.com/grafana/agent/component/otelcol/exporter/logging"                 // Import otelcol.exporter.logging
	_ "github.com/grafana/agent/component/otelcol/exporter/loki"                    // Import otelcol.exporter.loki
	_ "github.com/grafana/agent/component/otelcol/exporter/otlp"                    // Import otelcol.exporter.otlp
	_ "github.com/grafana/agent/component/otelcol/exporter/otlphttp"                // Import otelcol.exporter.otlphttp
	_ "github.com/grafana/agent/component/otelcol/exporter/prometheus"              // Import otelcol.exporter.prometheus
	_ "github.com/grafana/agent/component/otelcol/extension/jaeger_remote_sampling" // Import otelcol.extension.jaeger_remote_sampling
	_ "github.com/grafana/agent/component/otelcol/processor/attributes"             // Import otelcol.processor.attributes
	_ "github.com/grafana/agent/component/otelcol/processor/batch"                  // Import otelcol.processor.batch
	_ "github.com/grafana/agent/component/otelcol/processor/discovery"              // Import otelcol.processor.discovery
	_ "github.com/grafana/agent/component/otelcol/processor/memorylimiter"          // Import otelcol.processor.memory_limiter
	_ "github.com/grafana/agent/component/otelcol/processor/span"                   // Import otelcol.processor.span
	_ "github.com/grafana/agent/component/otelcol/processor/tail_sampling"          // Import otelcol.processor.tail_sampling
	_ "github.com/grafana/agent/component/otelcol/receiver/jaeger"                  // Import otelcol.receiver.jaeger
	_ "github.com/grafana/agent/component/otelcol/receiver/kafka"                   // Import otelcol.receiver.kafka
	_ "github.com/grafana/agent/component/otelcol/receiver/loki"                    // Import otelcol.receiver.loki
	_ "github.com/grafana/agent/component/otelcol/receiver/opencensus"              // Import otelcol.receiver.opencensus
	_ "github.com/grafana/agent/component/otelcol/receiver/otlp"                    // Import otelcol.receiver.otlp
	_ "github.com/grafana/agent/component/otelcol/receiver/prometheus"              // Import otelcol.receiver.prometheus
	_ "github.com/grafana/agent/component/otelcol/receiver/zipkin"                  // Import otelcol.receiver.zipkin
	_ "github.com/grafana/agent/component/prometheus/exporter/apache"               // Import prometheus.exporter.apache
	_ "github.com/grafana/agent/component/prometheus/exporter/blackbox"             // Import prometheus.exporter.blackbox
	_ "github.com/grafana/agent/component/prometheus/exporter/cloudwatch"           // Import prometheus.exporter.cloudwatch
	_ "github.com/grafana/agent/component/prometheus/exporter/consul"               // Import prometheus.exporter.consul
	_ "github.com/grafana/agent/component/prometheus/exporter/dnsmasq"              // Import prometheus.exporter.dnsmasq
	_ "github.com/grafana/agent/component/prometheus/exporter/elasticsearch"        // Import prometheus.exporter.elasticsearch
	_ "github.com/grafana/agent/component/prometheus/exporter/gcp"                  // Import prometheus.exporter.gcp
	_ "github.com/grafana/agent/component/prometheus/exporter/github"               // Import prometheus.exporter.github
	_ "github.com/grafana/agent/component/prometheus/exporter/kafka"                // Import prometheus.exporter.kafka
	_ "github.com/grafana/agent/component/prometheus/exporter/memcached"            // Import prometheus.exporter.memcached
	_ "github.com/grafana/agent/component/prometheus/exporter/mongodb"              // Import prometheus.exporter.mongodb
	_ "github.com/grafana/agent/component/prometheus/exporter/mssql"                // Import prometheus.exporter.mssql
	_ "github.com/grafana/agent/component/prometheus/exporter/mysql"                // Import prometheus.exporter.mysql
	_ "github.com/grafana/agent/component/prometheus/exporter/oracledb"             // Import prometheus.exporter.oracledb
	_ "github.com/grafana/agent/component/prometheus/exporter/postgres"             // Import prometheus.exporter.postgres
	_ "github.com/grafana/agent/component/prometheus/exporter/process"              // Import prometheus.exporter.process
	_ "github.com/grafana/agent/component/prometheus/exporter/redis"                // Import prometheus.exporter.redis
	_ "github.com/grafana/agent/component/prometheus/exporter/snmp"                 // Import prometheus.exporter.snmp
	_ "github.com/grafana/agent/component/prometheus/exporter/snowflake"            // Import prometheus.exporter.snowflake
	_ "github.com/grafana/agent/component/prometheus/exporter/squid"                // Import prometheus.exporter.squid
	_ "github.com/grafana/agent/component/prometheus/exporter/statsd"               // Import prometheus.exporter.statsd
	_ "github.com/grafana/agent/component/prometheus/exporter/unix"                 // Import prometheus.exporter.unix
	_ "github.com/grafana/agent/component/prometheus/exporter/windows"              // Import prometheus.exporter.windows
	_ "github.com/grafana/agent/component/prometheus/operator/podmonitors"          // Import prometheus.operator.podmonitors
	_ "github.com/grafana/agent/component/prometheus/operator/probes"               // Import prometheus.operator.probes
	_ "github.com/grafana/agent/component/prometheus/operator/servicemonitors"      // Import prometheus.operator.servicemonitors
	_ "github.com/grafana/agent/component/prometheus/receive_http"                  // Import prometheus.receive_http
	_ "github.com/grafana/agent/component/prometheus/relabel"                       // Import prometheus.relabel
	_ "github.com/grafana/agent/component/prometheus/remotewrite"                   // Import prometheus.remote_write
	_ "github.com/grafana/agent/component/prometheus/scrape"                        // Import prometheus.scrape
	_ "github.com/grafana/agent/component/pyroscope/ebpf"                           // Import pyroscope.ebpf
	_ "github.com/grafana/agent/component/pyroscope/scrape"                         // Import pyroscope.scrape
	_ "github.com/grafana/agent/component/pyroscope/write"                          // Import pyroscope.write
	_ "github.com/grafana/agent/component/remote/http"                              // Import remote.http
	_ "github.com/grafana/agent/component/remote/s3"                                // Import remote.s3
	_ "github.com/grafana/agent/component/remote/vault"                             // Import remote.vault
)
