package cmd

import (
	// homedir "github.com/mitchellh/go-homedir"
	"localserve/localserve/internal/tuned_log"

	"github.com/spf13/cobra"
)

var (
	tunedLogger = tuned_log.GetDefaultLogger() // logger should be closed in Execute()
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
func Execute() error {
	defer tuned_log.CloseDefaultLogger()
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
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
