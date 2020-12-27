package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "net/http"
  "log"
)

func init() {
  // default listening addr and port
  viper.SetDefault("listenAddr", "127.0.0.1")
  viper.SetDefault("listenPort", "3223")

  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "Start server",
  Long:  `Start server on your device`,
  Run: func(cmd *cobra.Command, args []string) {
    log.Fatal(startServer())
  },
}

func getListenAddr() string {
  listenAddr := viper.GetString("listenAddr")
  listenPort := viper.GetString("listenPort")
  return listenAddr + ":" + listenPort
}

func startServer() error {
  http.HandleFunc("/", indexHandler)

  fmt.Println("Starting server at", getListenAddr())

  return http.ListenAndServe(getListenAddr(), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("indexHandler was called\n"))
}
