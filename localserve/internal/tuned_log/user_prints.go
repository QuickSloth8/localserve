package tuned_log

import (
	"fmt"
	"log"
)

var (
	silent bool
)

func SetSilent(s bool) {
	silent = s
}

// logs info message, and prints it if silen == false
func InfoPrintToUser(msg string) {
	if silent == false {
		fmt.Println(msg)
	}
}

// logs error message, and prints it if silen == false
func ErrorPrintToUser(msg string) {
	if silent == false {
		fmt.Println("ERROR: ", msg)
	}
}

func Fatal(err error) {
	log.Fatal(err)
}
