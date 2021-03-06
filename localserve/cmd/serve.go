package cmd

import (
	"localserve/localserve/internal"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultServeRootHelp string
	currDirStr           string = "current directory - "

	defaultServePort   string
	defaultServeRoot   string
	defaultIdleTimeout int

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

	defaultIdleTimeout = 120

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

	serveCmd.PersistentFlags().Int(
		"auto-term",
		defaultIdleTimeout,
		"Auto terminate after duration in seconds\nSet it to zero or less to keep server running",
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
	viper.BindPFlag("auto-term",
		serveCmd.PersistentFlags().Lookup("auto-term"))

	viper.SetDefault("serveAddr", internal.GetIp())

	// add command
	rootCmd.AddCommand(serveCmd)
}
