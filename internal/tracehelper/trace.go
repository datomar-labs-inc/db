package tracehelper

import "go.opentelemetry.io/otel/trace"

var Tracer trace.Tracer

func SetTracer(t trace.Tracer) {
	Tracer = t
}
