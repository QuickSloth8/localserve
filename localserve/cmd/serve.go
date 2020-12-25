package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
  Use:   "serve",
  Short: "Start server",
  Long:  `Start server on your device`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Starting Server ...")
  },
}
