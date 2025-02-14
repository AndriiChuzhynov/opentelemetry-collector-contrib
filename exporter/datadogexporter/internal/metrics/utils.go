// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package metrics // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter/internal/metrics"

import (
	"fmt"

	"go.opentelemetry.io/collector/component"
	"gopkg.in/zorkian/go-datadog-api.v2"
)

type MetricDataType string

const (
	// Gauge is the Datadog Gauge metric type
	Gauge MetricDataType = "gauge"
	// Count is the Datadog Count metric type
	Count MetricDataType = "count"
)

// newMetric creates a new Datadog metric given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func newMetric(name string, ts uint64, value float64, tags []string) datadog.Metric {
	// Transform UnixNano timestamp into Unix timestamp
	// 1 second = 1e9 ns
	timestamp := float64(ts / 1e9)

	metric := datadog.Metric{
		Points: []datadog.DataPoint{[2]*float64{&timestamp, &value}},
		Tags:   tags,
	}
	metric.SetMetric(name)
	return metric
}

// NewMetric creates a new Datadog metric given a name, a type, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewMetric(name string, dt MetricDataType, ts uint64, value float64, tags []string) datadog.Metric {
	metric := newMetric(name, ts, value, tags)
	metric.SetType(string(dt))
	return metric
}

// NewGauge creates a new Datadog Gauge metric given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewGauge(name string, ts uint64, value float64, tags []string) datadog.Metric {
	return NewMetric(name, Gauge, ts, value, tags)
}

// NewCount creates a new Datadog count metric given a name, a Unix nanoseconds timestamp
// a value and a slice of tags
func NewCount(name string, ts uint64, value float64, tags []string) datadog.Metric {
	return NewMetric(name, Count, ts, value, tags)
}

// DefaultMetrics creates built-in metrics to report that an exporter is running
func DefaultMetrics(exporterType string, hostname string, timestamp uint64, buildInfo component.BuildInfo) []datadog.Metric {
	var tags []string

	if buildInfo.Version != "" {
		tags = append(tags, "version:"+buildInfo.Version)
	}

	if buildInfo.Command != "" {
		tags = append(tags, "command:"+buildInfo.Command)
	}

	metrics := []datadog.Metric{
		NewGauge(fmt.Sprintf("otel.datadog_exporter.%s.running", exporterType), timestamp, 1.0, tags),
	}

	for i := range metrics {
		metrics[i].SetHost(hostname)
	}

	return metrics
}
