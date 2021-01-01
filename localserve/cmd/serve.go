package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
  "net/http"
  "log"
  "os"
  "localserve/localserve/internal"
)



var (
  uninitializedServeRootFlag string = "current directory"
  defaultServeRoot string

  // flagServeAddr string
  flagServePort string
  flagServeRoot string

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
  // serveCmd.PersistentFlags().StringVar(
  //   &flagServeAddr,
  //   "serveAddr",
  //   "127.0.0.1",
  //   "The IP to listen on",
  // )

  serveCmd.PersistentFlags().StringVar(
    &flagServePort,
    "servePort",
    "3223",
    "The port to listen on",
  )

  serveCmd.PersistentFlags().StringVar(
    &flagServeRoot,
    "serveRoot",
    uninitializedServeRootFlag,
    "The directory to be served",
  )

  // viper.BindPFlag("serveAddr",
  //   serveCmd.PersistentFlags().Lookup("serveAddr"))
  viper.BindPFlag("servePort",
    serveCmd.PersistentFlags().Lookup("servePort"))
  viper.BindPFlag("serveRoot",
    serveCmd.PersistentFlags().Lookup("serveRoot"))

  // defaults
  var err error
  defaultServeRoot, err = os.Getwd()
  if err != nil {
    panic(err)
  }

  viper.SetDefault("serveAddr", internal.GetIp())

  // add command
  rootCmd.AddCommand(serveCmd)
}

func getFullServeAddr() string {
  serveAddr := viper.GetString("serveAddr")
  servePort := viper.GetString("servePort")
  return serveAddr + ":" + servePort
}

func getServeConfigsStr() string {
  strFullServeAddr := getFullServeAddr()
  strServeRoot := viper.GetString("serveRoot")

  return fmt.Sprintf("Serving %q at %q",
    strServeRoot, strFullServeAddr)
}

func startServer() error {
  serveRoot := viper.GetString("serveRoot")
  if serveRoot == uninitializedServeRootFlag {
    serveRoot = defaultServeRoot
    viper.Set("serveRoot", defaultServeRoot)
  }

  fs := http.FileServer(http.Dir(serveRoot))
  fmt.Println(getServeConfigsStr())
  return http.ListenAndServe(getFullServeAddr(), fs)
}
