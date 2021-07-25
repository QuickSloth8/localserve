package cmd

import (
	// homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	// cfgFile     string
	// userLicense string

	rootCmd = &cobra.Command{
		Use:   "localserve",
		Short: "Turn your device into a local server",
		Long: `This app enables you to serve and share files easily, by turning a
device to a local http server`,
	}
)

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	rootCmd.PersistentFlags().Bool(
		"log",
		false,
		"log for monitoring and debugging",
	)

	rootCmd.PersistentFlags().Bool(
		"log-to-file",
		false,
		`log output to file ./log.log`,
	)

	viper.BindPFlag("log",
		rootCmd.PersistentFlags().Lookup("log"))
	viper.BindPFlag("log-to-file",
		rootCmd.PersistentFlags().Lookup("log-to-file"))
}

// func er(msg interface{}) {
// 	fmt.Println("Error:", msg)
// 	os.Exit(1)
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			er(err)
// 		}
//
// 		// Search config in home directory with name ".cobra" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".cobra")
// 	}
//
// 	viper.AutomaticEnv()
//
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
