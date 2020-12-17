package main

import (
  "net/http"
  "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("REQ: %v\n", r.URL)
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("sdfsafsf\n"))
}

func main() {
  http.Handle("/", http.HandlerFunc(handler))
  addr := ":8000"
  log.Printf("starting server at %s\n", addr)
  log.Fatal(http.ListenAndServe(addr, nil))
}
