package unmerged

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bogdandrutu/metrics-proto/core"
	v1 "github.com/bogdandrutu/metrics-proto/unmerged/gen/common/v1"
	otlpmetric "github.com/bogdandrutu/metrics-proto/unmerged/gen/metrics/v1"
)

// Generator allows to generate a ExportRequest.
type Generator struct {
	random     *rand.Rand
	tracesSent uint64
	spansSent  uint64
}

func NewGenerator() *Generator {
	return &Generator{
		random: rand.New(rand.NewSource(99)),
	}
}

func (g *Generator) Name() string {
	return "unmerged"
}

func (g *Generator) NewMessage() core.ProtoMessage {
	return &otlpmetric.ResourceMetrics{}
}

func (g *Generator) GenerateGaugeMetrics(metricsPerBatch int) core.ProtoMessage {
	il := otlpmetric.InstrumentationLibraryMetrics{}
	for i := 0; i < metricsPerBatch; i++ {
		startTime := time.Date(2019, 10, 31, 10, 11, 12, 13, time.UTC)
		il.Metrics = append(il.Metrics, genInt64Gauge(startTime, i))
	}

	return &otlpmetric.ResourceMetrics{
		InstrumentationLibraryMetrics: []otlpmetric.InstrumentationLibraryMetrics{il},
	}
}

func genInt64Gauge(startTime time.Time, i int) otlpmetric.Metric {
	return otlpmetric.Metric{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Data: &otlpmetric.Metric_IntGauge{
			IntGauge: &otlpmetric.IntGauge{
				DataPoints: genGaugePoints(startTime, i),
			},
		},
	}
}

func genGaugePoints(startTime time.Time, i int) []otlpmetric.Int64DataPoint {
	var points []otlpmetric.Int64DataPoint
	for j := 0; j < 5; j++ {
		pointTs := startTime.Add(time.Duration(j) * time.Millisecond).UnixNano()
		point := otlpmetric.Int64DataPoint{
			Labels: []v1.StringKeyValue{
				{
					Key:   "key_1",
					Value: strconv.FormatInt(int64(i*10+j), 10),
				},
				{
					Key:   "key_2",
					Value: strconv.FormatInt(int64(i*10+j), 10),
				},
				{
					Key:   "key_3",
					Value: strconv.FormatInt(int64(i*10+j), 10),
				},
			},
			TimeUnixNano: uint64(pointTs),
			Value:        int64(i * j * 127),
		}
		points = append(points, point)
	}

	return points
}
