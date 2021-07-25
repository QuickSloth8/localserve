package tuned_log

import (
	"github.com/rs/zerolog"
)

type TunedLogger struct {
	zeroLogger zerolog.Logger
}

// TODO: write wrapper to check if tl is nil

func (tl *TunedLogger) Debug(msg string) {
	tl.zeroLogger.Debug().Msg(msg)
}

func (tl *TunedLogger) Info(msg string) {
	tl.zeroLogger.Info().Msg(msg)
}

func (tl *TunedLogger) Warn(msg string) {
	tl.zeroLogger.Warn().Msg(msg)
}

func (tl *TunedLogger) Error(msg string) {
	tl.zeroLogger.Error().Msg(msg)
}

func (tl *TunedLogger) Fatal(err error) {
	tl.zeroLogger.Fatal().Err(err)
}

func (tl *TunedLogger) Panic(err error) {
	tl.zeroLogger.Panic().Err(err)
}
