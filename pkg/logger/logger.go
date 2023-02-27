package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type (
	Logger struct {
		logger zerolog.Logger
	}
)

var (
	defaultLogger Logger
)

func init() {
	defaultLogger = Logger{log.Output(zerolog.ConsoleWriter{Out: os.Stderr})}
}

func Module(name string) Logger {
	log := defaultLogger.logger.With().Str("module", name).Logger()
	return Logger{log}
}

func (l Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}
