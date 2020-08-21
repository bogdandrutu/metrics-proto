package merged

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/bogdandrutu/metrics-proto/core"
	v1 "github.com/bogdandrutu/metrics-proto/merged/gen/common/v1"
	otlpmetric "github.com/bogdandrutu/metrics-proto/merged/gen/metrics/v1"
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
	return "merged"
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
		Data: &otlpmetric.Metric_Gauge{
			Gauge: &otlpmetric.Gauge{
				MeasurementValueType: otlpmetric.MeasurementValueType_MEASUREMENT_VALUE_TYPE_INT64,
				DataPoints:           genGaugePoints(startTime, i),
			},
		},
	}
}

func genGaugePoints(startTime time.Time, i int) []otlpmetric.ScalarDataPoint {
	var points []otlpmetric.ScalarDataPoint
	for j := 0; j < 5; j++ {
		pointTs := startTime.Add(time.Duration(j) * time.Millisecond).UnixNano()
		point := otlpmetric.ScalarDataPoint{
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
			Int64Value:   int64(i * j * 127),
		}
		points = append(points, point)
	}

	return points
}
