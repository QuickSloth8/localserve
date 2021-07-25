package tuned_log

import (
	"errors"
	"io"
	"os"
	"sync"

	"github.com/rs/zerolog"
)

// watches after defaultLogger usages, to insure cleaning
// only after all instances called CloseDefaultLogger()
type defaultAppLoggerRefsCount struct {
	numRefs int
	mux     sync.Mutex
}

// add one ref to the count
func (rc *defaultAppLoggerRefsCount) IncCount() {
	rc.mux.Lock()
	defer rc.mux.Unlock()
	rc.numRefs++
}

// remove one ref from count, and initiate cleaning process
// if no other references exist
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
	initOnce         = sync.Once{} // initDefaultLogger() once
	refsCount        = defaultAppLoggerRefsCount{}
	logFile          *os.File
	logging          bool
	logToFile        bool
	defaultAppLogger *TunedLogger
)

func initDefaultLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	var out io.Writer

	if logging {
		if logToFile {
			var err error
			// TODO: implement a log file size limitation or log rotation method
			out, err = os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				panic(err)
			}
		} else {
			out = zerolog.NewConsoleWriter()
		}

		// var err error
		// // TODO: implement a log file size limitation or log rotation method
		// out, err = os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		// if err != nil {
		// 	panic(err)
		// }

		defaultAppLogger = &TunedLogger{zerolog.New(out).With().Timestamp().Logger()}
	}
}

func destroyDefaultLogger() {
	refsCount.mux.Lock()
	defer refsCount.mux.Unlock()

	// close all log files
	if logFile != nil {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}

	// reset package state to a fresh one
	defaultAppLogger = nil // let garbage-collector clear previous instance
	refsCount.numRefs = 0  // reset instance count
	initOnce = sync.Once{} // new fresh instance
}

func GetDefaultLogger() *TunedLogger {
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
