package tuned_log

import (
	"errors"
	"os"
	"sync"

	"github.com/rs/zerolog"
	// "github.com/rs/zerolog/log"
)

type defaultAppLoggerRefsCount struct {
	numRefs int
	mux     sync.Mutex
}

func (rc *defaultAppLoggerRefsCount) IncCount() {
	rc.mux.Lock()
	defer rc.mux.Unlock()

	rc.numRefs++
}

func (rc *defaultAppLoggerRefsCount) DecCount() (empty bool, err error) {
	rc.mux.Lock()
	defer rc.mux.Unlock()

	rc.numRefs--
	if rc.numRefs == 0 {
		empty = true
	} else if rc.numRefs < 0 {
		err = errors.New("unexpected state! count is negative!!")
	}
	return empty, err
}

var (
	initOnce         = sync.Once{} // initDefaultLogger() once, if non-zero val exists
	refsCount        = defaultAppLoggerRefsCount{}
	defaultAppLogger *defaultLogger
)

func initDefaultLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defaultAppLogger = &defaultLogger{zerolog.New(logFile).With().Timestamp().Logger()}
}

func destroyDefaultLogger() {
	refsCount.mux.Lock()
	defer refsCount.mux.Unlock()

	// close any files if they exist

	defaultAppLogger = nil // let garbage-collector clear previous instance
	refsCount.numRefs = 0  // reset instance count
	initOnce = sync.Once{} // new fresh instance
}

func GetDefaultLogger() *defaultLogger {
	initOnce.Do(initDefaultLogger)
	refsCount.IncCount()
	return defaultAppLogger
}

func CloseDefaultLogger() {
	empty, err := refsCount.DecCount()
	if err != nil {
		panic(err)
	}

	if empty {
		destroyDefaultLogger()
	}
}
