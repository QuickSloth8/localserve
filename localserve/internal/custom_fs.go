package internal

import (
	"fmt"
	"localserve/localserve/internal/tuned_log"
	"net/http"
)

// Custom Handler that prints requested URLs before serving
type CustomFileServer struct {
	Handler http.Handler
}

func (cfs CustomFileServer) PrintRequestSummary(req *http.Request) {
	msg := fmt.Sprintf("%s: %s", req.Method, req.URL)
	fmt.Println(msg) // this should be called only when a specific flag is present
	tuned_log.InfoOnce(msg)
}

func (cfs CustomFileServer) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	cfs.PrintRequestSummary(req)
	cfs.Handler.ServeHTTP(respW, req)
}
