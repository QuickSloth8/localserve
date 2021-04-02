package tuned_log

import (
	"fmt"
)

func SetSilent(s bool) {
	silent = s
}

func PrintInfoToUser(msg string, logger *defaultLogger) {
	logger.Info(msg)
	if silent == false {
		fmt.Println(msg)
	}
}

func PrintErrorToUser(msg string, logger *defaultLogger) {
	logger.Error(msg)
	if silent == false {
		fmt.Println("ERROR: ", msg)
	}
}
