package tuned_log

import (
	"fmt"
)

func SetSilent(s bool) {
	silent = s
}

// logs info message, and prints it if silen == false
func InfoPrintToUser(msg string, logger *defaultLogger) {
	logger.Info(msg)
	if silent == false {
		fmt.Println(msg)
	}
}

// logs error message, and prints it if silen == false
func ErrorPrintToUser(msg string, logger *defaultLogger) {
	logger.Error(msg)
	if silent == false {
		fmt.Println("ERROR: ", msg)
	}
}
