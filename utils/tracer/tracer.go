package tracer

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

//InitTracer ...
func InitTracer() *tracesdk.TracerProvider {
	tp, err := tracerProvider()
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	return tp
}

func tracerProvider() (*tracesdk.TracerProvider, error) {
	traceProvider, err := jaeger.InstallNewPipeline(
		jaeger.WithCollectorEndpoint(
			jaeger.WithEndpoint("http://localhost:14268/api/traces"),
		),
	)
	return traceProvider, err
}
