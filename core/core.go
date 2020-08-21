package core

import (
	"github.com/gogo/protobuf/proto"
)

type ProtoMessage interface {
	proto.Marshaler
	proto.Unmarshaler
}

type MetricGenerator interface {
	Name() string
	NewMessage() ProtoMessage
	GenerateGaugeMetrics(metricsPerBatch int) ProtoMessage
}
