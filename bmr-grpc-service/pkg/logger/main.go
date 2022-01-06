package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func Instance() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return logger
}
