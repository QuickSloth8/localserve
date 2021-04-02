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

func getFileSystem(doneChan chan os.Signal) *internal.CustomFileServer {
	flagIdleTimeout := viper.GetInt("auto-term")

	handleServeRootCleaning() // removes currDirStr from serveRoot
	serveRoot := viper.GetString("serveRoot")
	handler := http.FileServer(http.Dir(serveRoot))

	if flagIdleTimeout > 0 {
		convertedIdleTimeout := time.Duration(flagIdleTimeout) * time.Second
		return internal.NewCustomFileServerWithTimeout(
			handler,
			convertedIdleTimeout,
			doneChan,
		)
	} else {
		return internal.NewCustomFileServer(handler)
	}
}

func startServer() {
	defer tuned_log.InfoPrintToUser("\nThank you for choosing LocalServe :)\n", tunedLogger)

	// set global silent output flag in tuned_log package
	tuned_log.SetSilent(viper.GetBool("silent"))

	// set keyboard interrupt listener channel
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	fs := getFileSystem(done)

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

	timeoutSecs := 30 * time.Second
	msg := fmt.Sprintf("Server termination initiated (%v max)", timeoutSecs)
	tuned_log.InfoPrintToUser(msg, tunedLogger)

	ctx, cancel := context.WithTimeout(context.Background(), timeoutSecs)
	// ctx := context.Background()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown Failed - ", err)
	}
	defer cancel()
}
