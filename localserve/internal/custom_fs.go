package internal

import (
	"fmt"
	"localserve/localserve/internal/tuned_log"
	"net/http"
)

// Custom Handler that prints requested URLs before serving
type CustomFileServer struct {
	Handler http.Handler
	Silent  bool
}

func (cfs CustomFileServer) PrintRequestSummary(req *http.Request) {
	tunedLogger := tuned_log.GetDefaultLogger()
	defer tuned_log.CloseDefaultLogger()
	msg := fmt.Sprintf("%s %s", req.Method, req.URL)
	tuned_log.PrintInfoToUser(msg, tunedLogger, cfs.Silent)
}

func (cfs CustomFileServer) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	cfs.PrintRequestSummary(req)
	cfs.Handler.ServeHTTP(respW, req)
}
