package tuned_log

import (
	"fmt"
)

func SetLogging(l bool) {
	fmt.Println("######### Setting logging to", l)
	logging = l
}

func SetLogToFile(f bool) {
	logToFile = f
}

// logs info message, and prints it if silen == false
func (logger *TunedLogger) InfoPrintToUser(msg string) {
	if logger != nil {
		logger.Info(msg)
	}
	fmt.Println(msg)
}

// logs error message, and prints it if silen == false
func (logger *TunedLogger) ErrorPrintToUser(msg string) {
	if logger != nil {
		logger.Error(msg)
	}
	fmt.Println("ERROR: ", msg)
}
