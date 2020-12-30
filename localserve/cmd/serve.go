package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "net/http"
  "log"
  "os"
)



var (
  listenAddr string
  listenPort string

  serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Start server",
    Long:  `Start server on your device`,
    Run: func(cmd *cobra.Command, args []string) {
      log.Fatal(startServer())
    },
  }
)

func init() {
  // serveCmd.PersistentFlags().StringVar(&listenAddr, "listenAddr", "127.0.0.1", "The IP to listen on")
  serveCmd.PersistentFlags().StringVar(&listenPort, "listenPort", "3223", "The port to listen on")
  // viper.BindPFlag("listenAddr", serveCmd.PersistentFlags().Lookup("listenAddr"))
  viper.BindPFlag("listenPort", serveCmd.PersistentFlags().Lookup("listenPort"))

  serveRoot, err := os.Getwd()
  if err != nil {
    panic(err)
  }
  viper.SetDefault("serveRoot", serveRoot)

  rootCmd.AddCommand(serveCmd)
}

func getListenAddr() string {
  listenAddr := viper.GetString("listenAddr")
  listenPort := viper.GetString("listenPort")
  return listenAddr + ":" + listenPort
}

func startServer() error {
  serveRoot := viper.GetString("serveRoot")
  fs := http.FileServer(http.Dir(serveRoot))

  fmt.Println("Starting server at", getListenAddr())
  fmt.Println("Serving files in directory", serveRoot)
  return http.ListenAndServe(getListenAddr(), fs)
}
