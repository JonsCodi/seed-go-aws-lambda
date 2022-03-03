package cmd

import (
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use: "start",
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringP("env-file", "e", "", "config file")
}
