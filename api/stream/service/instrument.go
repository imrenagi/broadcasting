package services

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/metric/unit"
)

var tracer = otel.Tracer("github.com/imrenagi/broadcasting/api/stream/services")

var meter = global.Meter("github.com/imrenagi/broadcasting/api/stream/services")

var streamCounter, _ = meter.SyncInt64().Counter("stream_count",
	instrument.WithDescription("number of stream "),
	instrument.WithUnit(unit.Dimensionless))

