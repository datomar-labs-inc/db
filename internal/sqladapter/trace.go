package sqladapter

import "go.opentelemetry.io/otel/trace"

var tracer trace.Tracer

func SetTracer(t trace.Tracer) {
	tracer = t
}
