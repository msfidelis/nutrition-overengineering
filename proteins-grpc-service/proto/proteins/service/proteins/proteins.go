package proteins

import (
	calculator "proteins-grpc-service/pkg/proteins"

	"proteins-grpc-service/pkg/logger"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()
	tracer := otel.Tracer("proteins-grpc-server")
	_, span := tracer.Start(ctx, "SayHello")
	defer span.End()

	log.Info().
		Str("Service", "proteins").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Msg("Calculating Proteins Necessity")

	recommendation := calculator.Calc(int64(in.Weight))

	log.Info().
		Str("Service", "proteins").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Int64("Proteins Necessity", recommendation).
		Msg("Proteins Necessity calculated")

	return &Response{Value: recommendation, Unit: "kcal"}, nil
}
