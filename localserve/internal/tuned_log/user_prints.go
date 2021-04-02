package tuned_log

import (
	"fmt"
)

func PrintInfoToUser(msg string, logger *defaultLogger, silent bool) {
	logger.Info(msg)
	if silent == false {
		fmt.Println(msg)
	}
}

func PrintErrorToUser(msg string, logger *defaultLogger, silent bool) {
	logger.Error(msg)
	if silent == false {
		fmt.Println("ERROR: ", msg)
	}
}
