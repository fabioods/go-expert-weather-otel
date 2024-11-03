package webserver

import (
	"context"
	"net/http"

	"github.com/fabioods/go-expert-wheater-lab/pkg/otel"
	"go.opentelemetry.io/otel/trace"
)

type TraceMiddleware struct {
	tracer trace.Tracer
}

func NewTraceMiddleware(tracer trace.Tracer) *TraceMiddleware {
	return &TraceMiddleware{tracer: tracer}
}

func (tm *TraceMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tm.tracer.Start(r.Context(), r.Method+" "+r.URL.Path)
		defer span.End()

		ctx = context.WithValue(ctx, otel.TracerKey, tm.tracer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
