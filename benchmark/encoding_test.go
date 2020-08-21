package benchmark

import (
	"testing"

	"github.com/bogdandrutu/metrics-proto/core"
	"github.com/bogdandrutu/metrics-proto/merged"
	"github.com/bogdandrutu/metrics-proto/mergedfixed"
	"github.com/bogdandrutu/metrics-proto/unmerged"
	"github.com/bogdandrutu/metrics-proto/unmergedfixed"

	"github.com/stretchr/testify/assert"
)

var tests = []core.MetricGenerator{
	merged.NewGenerator(),
	mergedfixed.NewGenerator(),
	unmerged.NewGenerator(),
	unmergedfixed.NewGenerator(),
}

func BenchmarkEncodeDecode(b *testing.B) {
	for _, v := range tests {
		msg := v.GenerateGaugeMetrics(10)
		b.Run(v.Name(), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				buf, err := msg.Marshal()
				assert.NoError(b, err)
				nMsg := v.NewMessage()
				assert.NoError(b, nMsg.Unmarshal(buf))
			}
		})
	}
}
