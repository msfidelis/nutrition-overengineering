package water

import (
	calculator "water-grpc-service/pkg/water"

	"water-grpc-service/pkg/logger"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()
	tracer := otel.Tracer("water-grpc-server")
	_, span := tracer.Start(ctx, "SayHello")
	defer span.End()

	log.Info().
		Str("Service", "water").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Msg("Calculating Recommended Water Consume")

	recommendation := calculator.Calc(in.Weight)

	log.Info().
		Str("Service", "water").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Float64("Value", recommendation).
		Str("Unit", "ml").
		Msg("Water Consume Recommended")

	return &Response{Value: recommendation, Unit: "ml"}, nil
}
