package internal

import (
	"fmt"
	"localserve/localserve/internal/tuned_log"
	"net/http"
	"os"
	"syscall"
	"time"
)

// Custom Handler that prints requested URLs before serving
type CustomFileServer struct {
	Handler  http.Handler
	atw      *AutoTerminateWatch // for auto-termination
	termChan chan os.Signal
}

func NewCustomFileServerWithTimeout(h http.Handler, fsTermChan chan os.Signal, maxIdleTime time.Duration,
	atwTermChan chan os.Signal) *CustomFileServer {

	atw := &AutoTerminateWatch{
		MaxIdleTimeout: maxIdleTime,
		currentTime:    maxIdleTime,
		termChan:       atwTermChan,
	}

	cfs := &CustomFileServer{
		Handler:  h,
		atw:      atw,
		termChan: fsTermChan,
	}

	cfs.atw.StartTimerOnce()

	return cfs
}

func NewCustomFileServer(h http.Handler, fsTermChan chan os.Signal) *CustomFileServer {
	return &CustomFileServer{
		Handler: h, atw: nil,
		termChan: fsTermChan,
	}
}

func (cfs CustomFileServer) PrintRequestSummary(req *http.Request) {
	tunedLogger := tuned_log.GetDefaultLogger()
	defer tuned_log.CloseDefaultLogger()
	msg := fmt.Sprintf("%s %s", req.Method, req.URL)
	tuned_log.InfoPrintToUser(msg, tunedLogger)
}

func (cfs CustomFileServer) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" && req.URL.String() == "/shutdown" {
		tunedLogger := tuned_log.GetDefaultLogger()
		tuned_log.InfoPrintToUser("Remote shutdown requested", tunedLogger)
		tuned_log.CloseDefaultLogger()

		cfs.termChan <- syscall.SIGTERM
		return
	}

	if cfs.atw != nil {
		cfs.atw.ResetTimer() // every new request, the atw timer gets reset
	}

	cfs.PrintRequestSummary(req)
	cfs.Handler.ServeHTTP(respW, req)
}
