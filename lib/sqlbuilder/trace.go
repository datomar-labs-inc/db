package sqlbuilder

import (
	"github.com/datomar-labs-inc/db/internal/sqladapter"
	"go.opentelemetry.io/otel/trace"
)

func SetTracer(t trace.Tracer) {
	sqladapter.SetTracer(t)
}
