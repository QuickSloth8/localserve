package tuned_log

import (
	"fmt"
)

func SetSilent(s bool) {
	silent = s
}

func InfoPrintToUser(msg string, logger *defaultLogger) {
	logger.Info(msg)
	if silent == false {
		fmt.Println(msg)
	}
}

func ErrorPrintToUser(msg string, logger *defaultLogger) {
	logger.Error(msg)
	if silent == false {
		fmt.Println("ERROR: ", msg)
	}
}
