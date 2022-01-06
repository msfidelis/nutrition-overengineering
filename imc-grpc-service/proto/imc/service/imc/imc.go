package imc

import (
	calculator "imc-grpc-service/pkg/imc"

	"imc-grpc-service/pkg/logger"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Response, error) {
	log := logger.Instance()

	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Msg("Calculating imc")

	imcCalc, class := calculator.Calc(in.Weight, in.Height)

	log.Info().
		Float64("Weight", in.Weight).
		Float64("Height", in.Height).
		Float64("IMC", imcCalc).
		Str("Class", class).
		Msg("IMC Calculated")

	return &Response{Imc: imcCalc, Class: class}, nil
}
