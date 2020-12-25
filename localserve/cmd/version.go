package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version",
  Long:  `Print the version of the current LocalServe app`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("LocalServe v0.1")
  },
}
