package cmd

import (
	"fmt"
	"localserve/localserve/internal"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultServeRootHelp string
	currDirStr           string = "current directory - "

	defaultServePort string
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
	// defaults
	defaultServePort = "3223"

	var err error
	defaultServeRoot, err = os.Getwd()
	if err != nil {
		panic(err)
	}

	defaultServeRootHelp = currDirStr + defaultServeRoot

	// serveCmd.PersistentFlags().StringVar(
	//   &flagServeAddr,
	//   "serveAddr",
	//   "127.0.0.1",
	//   "The IP to listen on",
	// )

	serveCmd.PersistentFlags().StringVar(
		&flagServePort,
		"servePort",
		defaultServePort,
		"The port to listen on",
	)

	serveCmd.PersistentFlags().StringVar(
		&flagServeRoot,
		"serveRoot",
		defaultServeRootHelp,
		"The directory to be served",
	)

	// bind command flags to viper
	// viper.BindPFlag("serveAddr",
	//   serveCmd.PersistentFlags().Lookup("serveAddr"))
	viper.BindPFlag("servePort",
		serveCmd.PersistentFlags().Lookup("servePort"))
	viper.BindPFlag("serveRoot",
		serveCmd.PersistentFlags().Lookup("serveRoot"))

	fmt.Println(internal.GetIp())
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

func handleServeRootCleaning() {
	currServeRoot := viper.GetString("serveRoot")
	if currServeRoot == defaultServeRootHelp {
		viper.Set("serveRoot", strings.Replace(currServeRoot, currDirStr, "", 1))
	}
}

func startServer() error {
	handleServeRootCleaning()
	serveRoot := viper.GetString("serveRoot")

	fs := http.FileServer(http.Dir(serveRoot))
	fmt.Println(getServeConfigsStr())
	return http.ListenAndServe(getFullServeAddr(), fs)
}
