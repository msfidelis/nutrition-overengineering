package tracer

import (
	"context"

	"go.opentelemetry.io/otel"

	// stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer() *sdktrace.TracerProvider {
	// exp, err := jaeger.New(
	// 	jaeger.WithCollectorEndpoint(
	// 		jaeger.WithEndpoint(
	// 			os.Getenv("JAEGER_COLLECTOR_ENDPOINT"),
	// 		),
	// 	),
	// )
	// if err != nil {
	// 	fmt.Println("Failed to init jaeger", err)
	// }
	ctx := context.Background()
	exp, err := otlptracehttp.New(ctx)
	if err != nil {
		panic(err)
	}
	tracerProvider := trace.NewTracerProvider(trace.WithBatcher(exp))
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()

	// tp := sdktrace.NewTracerProvider(
	// 	sdktrace.WithBatcher(exp),
	// 	sdktrace.WithSampler(sdktrace.AlwaysSample()),
	// 	sdktrace.WithResource(resource.NewWithAttributes(
	// 		semconv.SchemaURL,
	// 		semconv.ServiceNameKey.String("health-api"),
	// 	)),
	// )

	// otel.SetTextMapPropagator(
	// 	propagation.NewCompositeTextMapPropagator(
	// 		propagation.TraceContext{},
	// 		propagation.Baggage{},
	// 	),
	// )

	otel.SetTracerProvider(tracerProvider)

	return tracerProvider
}
