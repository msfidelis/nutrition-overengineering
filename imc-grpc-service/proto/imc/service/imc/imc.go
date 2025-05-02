package imc

import (
	calculator "imc-grpc-service/pkg/imc"

	"imc-grpc-service/pkg/logger"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()
	tracer := otel.Tracer("imc-grpc-server")
	_, span := tracer.Start(ctx, "SayHello")
	defer span.End()

	log.Info().
		Str("Service", "imc").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Msg("Calculating imc")

	imcCalc, class := calculator.Calc(in.Weight, in.Height)

	log.Info().
		Str("Service", "imc").
		Str("traceID", span.SpanContext().TraceID().String()).
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Float64("IMC", imcCalc).
		Str("Class", class).
		Msg("IMC Calculated")

	return &Response{Imc: imcCalc, Class: class}, nil
}
