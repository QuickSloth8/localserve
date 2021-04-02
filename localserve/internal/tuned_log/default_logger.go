package tuned_log

import (
	"github.com/rs/zerolog"
)

type defaultLogger struct {
	zeroLogger zerolog.Logger
}

func (dl defaultLogger) Debug(msg string) {
	dl.zeroLogger.Debug().Msg(msg)
}

func (dl defaultLogger) Info(msg string) {
	dl.zeroLogger.Info().Msg(msg)
}

func (dl defaultLogger) Warn(msg string) {
	dl.zeroLogger.Warn().Msg(msg)
}

func (dl defaultLogger) Error(msg string) {
	dl.zeroLogger.Error().Msg(msg)
}

func (dl defaultLogger) Fatal(msg string) {
	dl.zeroLogger.Fatal().Msg(msg)
}

func (dl defaultLogger) Panic(msg string) {
	dl.zeroLogger.Panic().Msg(msg)
}
