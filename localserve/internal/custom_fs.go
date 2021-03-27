package internal

import (
	"fmt"
	"net/http"
)

// Custom Handler that prints requested URLs before serving
type CustomFileServer struct {
	Handler http.Handler
}

func (cfs CustomFileServer) PrintRequestSummary(req *http.Request) {
	fmt.Printf("%s: %s\n", req.Method, req.URL)
}

func (cfs CustomFileServer) ServeHTTP(respW http.ResponseWriter, req *http.Request) {
	cfs.PrintRequestSummary(req)
	cfs.Handler.ServeHTTP(respW, req)
}
