package cmd

import (
	"context"
	"fmt"
	"localserve/localserve/internal"
	"localserve/localserve/internal/tuned_log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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
			startServer()
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
		"the port to listen on",
	)

	serveCmd.PersistentFlags().StringVar(
		&flagServeRoot,
		"serveRoot",
		defaultServeRootHelp,
		"the directory to be served",
	)

	serveCmd.PersistentFlags().Bool(
		"silent",
		false,
		"suppress all output to stdout",
	)

	// bind command flags to viper
	// viper.BindPFlag("serveAddr",
	//   serveCmd.PersistentFlags().Lookup("serveAddr"))
	viper.BindPFlag("servePort",
		serveCmd.PersistentFlags().Lookup("servePort"))
	viper.BindPFlag("serveRoot",
		serveCmd.PersistentFlags().Lookup("serveRoot"))
	viper.BindPFlag("silent",
		serveCmd.PersistentFlags().Lookup("silent"))

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

	return fmt.Sprintf("\nServing %q at %q\n",
		strServeRoot, strFullServeAddr)
}

func handleServeRootCleaning() {
	currServeRoot := viper.GetString("serveRoot")
	if currServeRoot == defaultServeRootHelp {
		viper.Set("serveRoot", strings.Replace(currServeRoot, currDirStr, "", 1))
	}
}

func startServer() error {
	defer tuned_log.InfoPrintToUser("\nThank you for choosing LocalServe :)\n", tunedLogger)

	// set global silent output flag in tuned_log package
	tuned_log.SetSilent(viper.GetBool("silent"))

	handleServeRootCleaning() // removes currDirStr from serveRoot
	serveRoot := viper.GetString("serveRoot")

	fs := internal.CustomFileServer{
		Handler: http.FileServer(http.Dir(serveRoot)),
	}

	tuned_log.InfoPrintToUser(getServeConfigsStr(), tunedLogger)
	srv := &http.Server{
		Addr:    getFullServeAddr(),
		Handler: fs,
	}

	// set keyboard interrupt listener channel
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	// start server
	go func() {
		defer func() {
			done <- syscall.SIGINT
		}()

		err := srv.ListenAndServe()
		if err != nil {
			// if port is already taken
			if _, ok := err.(*net.OpError); ok && err.(*net.OpError).Op == "listen" {
				// fmt.Printf("Opps! ... %q seems to be taken !\n\n", getFullServeAddr())
				msg := fmt.Sprintf("Opps! ... %q seems to be taken !\n\n", getFullServeAddr())
				tuned_log.ErrorPrintToUser(msg, tunedLogger)
			} else {
				tunedLogger.Fatal(err)
			}
		}
	}()

	// handle keyboard interrupt & graceful termination
	<-done

	timeoutSecs := 30 * time.Second
	msg := fmt.Sprintf("Server termination initiated (%v max)", timeoutSecs)
	tuned_log.InfoPrintToUser(msg, tunedLogger)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecs)
	// ctx := context.Background()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Failed - ", err)
	}
	defer cancel()

	return nil
}
