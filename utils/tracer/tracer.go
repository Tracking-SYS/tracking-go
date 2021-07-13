package tracer

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/trace/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
)

//InitTracer ...
func InitTracer(serviceName string) func() {
	endpoint := os.Getenv("JAEGER_ENDPOINT")
	if endpoint == "" {
		return func() {
			//Return shutdown handler function
		}
	}

	var endpointOpt jaeger.EndpointOption
	if strings.HasPrefix(endpoint, "http") {
		// using thrift protocol
		endpointOpt = jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint))
	} else {
		// jaeger agent endpoint (:6831)
		parts := strings.Split(endpoint, ":")
		endpointOpt = jaeger.WithAgentEndpoint(jaeger.WithAgentHost(parts[0]), jaeger.WithAgentPort(parts[0]))
	}

	withTags := []attribute.KeyValue{}
	withTags = append(withTags, semconv.ServiceNameKey.String(serviceName))
	exp, err := jaeger.NewRawExporter(endpointOpt)
	if err != nil {
		return nil
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(0.1)),
		sdktrace.WithResource(resource.NewWithAttributes(withTags...)),
	)
	otel.SetTracerProvider(tp)
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			fmt.Println("shutdown jaeger")
		}
	}
}
