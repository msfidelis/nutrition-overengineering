package calories

import (
	"calories-grpc-service/pkg/calories"
	"calories-grpc-service/pkg/logger"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {

	log := logger.Instance()
	tracer := otel.Tracer("calories-grpc-server")
	_, span := tracer.Start(ctx, "SayHello")
	defer span.End()

	log.Info().
		Float64("Necessity", in.Necessity).
		Msg("Calculating calories necessity")

	gain := calories.Gain(in.Necessity)
	loss := calories.Loss(in.Necessity)
	maintain := calories.Maintain(in.Necessity)

	log.Info().
		Str("Service", "calories").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Necessity", in.Necessity).
		Float64("Gain", gain).
		Float64("Loss", loss).
		Float64("Maintain", maintain).
		Msg("Daily calories necessity calculated")

	return &Response{Gain: gain, Loss: loss, Maintain: maintain}, nil
}
