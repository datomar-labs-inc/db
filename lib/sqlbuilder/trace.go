package sqlbuilder

import (
	"github.com/datomar-labs-inc/db/internal/tracehelper"
	"go.opentelemetry.io/otel/trace"
)

func SetTracer(t trace.Tracer) {
	tracehelper.SetTracer(t)
}
