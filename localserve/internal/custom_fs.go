package internal

import (
	"fmt"
	"localserve/localserve/internal/tuned_log"
	"net/http"
	"os"
	"time"
)

// Custom Handler that prints requested URLs before serving
type CustomFileServer struct {
	Handler http.Handler
	atw     *AutoTerminateWatch // for auto-termination
}

func NewCustomFileServerWithTimeout(h http.Handler, maxIdleTime time.Duration,
	termChannel chan os.Signal) *CustomFileServer {

	atw := &AutoTerminateWatch{
		MaxIdleTimeout: maxIdleTime,
		currentTime:    maxIdleTime,
		TermChan:       termChannel,
	}

	cfs := &CustomFileServer{
		Handler: h,
		atw:     atw,
	}

	cfs.atw.StartTimerOnce()

	return cfs
}

func (cfs CustomFileServer) PrintRequestSummary(req *http.Request) {
	tunedLogger := tuned_log.GetDefaultLogger()
	defer tuned_log.CloseDefaultLogger()
	msg := fmt.Sprintf("%s %s", req.Method, req.URL)
	tuned_log.InfoPrintToUser(msg, tunedLogger)
}

func (cfs CustomFileServer) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	cfs.atw.ResetTimer() // every new request, the atw timer gets reset
	cfs.PrintRequestSummary(req)
	cfs.Handler.ServeHTTP(respW, req)
}
