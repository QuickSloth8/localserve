package internal

import (
	"logger"
)

// STUFF TO THINK ABOUT INSIDE THIS PACKAGE
// is it concurrency safe? what happens if more than one routine/thread tries to use AppLogger
// should AppLogger be exported? (start with a capital letter)
// how to differ between returning a copy or a pointer or a reference?
// is there a better more elegant way to achive the goal? maybe use existing packages?
// should close logger crash on illegal call?

const maxCapacity = 10

var (
	loggerChan = make(chan bool, maxCapacity) // 10 goroutines can have reference to
	// 									the logger simultaneously (by calling getLogger())
	AppLogger logger.Logger // a sigleton logger instance
)

func getLogger() logger.Logger {
	// check if maxCapacity reached
	if len(loggerChan) >= maxCapacity {
		return nil
	}

	if AppLogger == nil {
		AppLogger = logger.Default()
	}

	loggerChan <- true
	return AppLogger
}

func closeLogger() {
	if _, ok := <-loggerChan; ok == false {
		terminateLogger()
		AppLogger.Fatal("closeLogger called before calling getLogger()")
	}
	if len(loggerChan) == 0 {
		terminateLogger()
	}
}

func terminateLogger() {
	// perform logger final termination, like closing the log file
}
