package bmr

import (
	calculator "bmr-grpc-service/pkg/bmr"

	"bmr-grpc-service/pkg/logger"

	"go.opentelemetry.io/otel"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()

	tracer := otel.Tracer("bmr-grpc-server")
	_, span := tracer.Start(ctx, "SayHello")
	defer span.End()

	log.Info().
		Str("Gender", in.Gender).
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Int64("Age", in.Age).
		Str("Activiry Rate", in.Activity).
		Msg("Calculating BMR and Calories Necessity")

	bmrCalc, necessity := calculator.Calc(in.Gender, in.Weight, in.Height, in.Age, in.Activity)

	log.Info().
		Str("Gender", in.Gender).
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Int64("Age", in.Age).
		Str("Activiry Rate", in.Activity).
		Float64("BMR", bmrCalc).
		Float64("Calorical Necessity", necessity).
		Msg("Basal Metabolical Rate Calculated")

	return &Response{Bmr: bmrCalc, Necessity: necessity}, nil
}
