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

	"github.com/spf13/viper"
)

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

func startServer() {
	defer tuned_log.InfoPrintToUser("\nThank you for choosing LocalServe :)\n", tunedLogger)

	// set global silent output flag in tuned_log package
	tuned_log.SetSilent(viper.GetBool("silent"))

	handleServeRootCleaning() // removes currDirStr from serveRoot
	serveRoot := viper.GetString("serveRoot")

	flagIdleTimeout := 10 * time.Second

	// set keyboard interrupt listener channel
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fs := internal.NewCustomFileServerWithTimeout(
		http.FileServer(http.Dir(serveRoot)),
		flagIdleTimeout,
		done,
	)
	// fs := internal.CustomFileServer{
	// 	Handler:     http.FileServer(http.Dir(serveRoot)),
	// 	MaxIdleTime: time.Duration(flagIdleTimeout),
	// 	Atw:         atw,
	// }

	tuned_log.InfoPrintToUser(getServeConfigsStr(), tunedLogger)
	srv := &http.Server{
		Addr:        getFullServeAddr(),
		Handler:     fs,
		IdleTimeout: time.Duration(5 * time.Second),
	}

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

	timeoutSecs := 3 * time.Second
	msg := fmt.Sprintf("Server termination initiated (%v max)", timeoutSecs)
	tuned_log.InfoPrintToUser(msg, tunedLogger)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecs)
	// ctx := context.Background()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Failed - ", err)
	}
	defer cancel()
}
