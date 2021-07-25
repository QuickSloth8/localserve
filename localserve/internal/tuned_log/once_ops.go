package tuned_log

// func DebugOnce(msg string) {
// 	onceLogger := GetDefaultLogger()
// 	defer CloseDefaultLogger()
// 	onceLogger.Debug(msg)
// }

func InfoOnce(msg string) {
	onceLogger := GetDefaultLogger()
	if onceLogger != nil {
		defer CloseDefaultLogger()
		onceLogger.Info(msg)
	}
}

func InfoPrintToUserOnce(msg string) {
	onceLogger := GetDefaultLogger()
	if onceLogger != nil {
		defer CloseDefaultLogger()
		onceLogger.InfoPrintToUser(msg)
	}
}

// func WarnOnce(msg string) {
// 	onceLogger := GetDefaultLogger()
// 	defer CloseDefaultLogger()
// 	onceLogger.Warn(msg)
// }

// func ErrorOnce(msg string) {
// 	onceLogger := GetDefaultLogger()
// 	defer CloseDefaultLogger()
// 	onceLogger.Error(msg)
// }

// func FatalOnce(err error) {
// 	onceLogger := GetDefaultLogger()
// 	defer CloseDefaultLogger()
// 	onceLogger.Fatal(err)
// }

// func PanicOnce(err error) {
// 	onceLogger := GetDefaultLogger()
// 	defer CloseDefaultLogger()
// 	onceLogger.Panic(err)
// }
