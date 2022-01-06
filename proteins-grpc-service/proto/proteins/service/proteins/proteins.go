package proteins

import (
	calculator "proteins-grpc-service/pkg/proteins"

	"proteins-grpc-service/pkg/logger"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()

	log.Info().
		Float64("Weight", in.Weight).
		Msg("Calculating Proteins Necessity")

	recommendation := calculator.Calc(int64(in.Weight))

	// Water Service Call

	// log.Info().
	// 	Float64("Weight", in.Weight).
	// 	Float64("Height", in.Height).
	// 	Float64("IMC", imcCalc).
	// 	Str("Class", class).
	// 	Msg("IMC Calculated")

	return &Response{Value: recommendation, Unit: "kcal"}, nil
}
