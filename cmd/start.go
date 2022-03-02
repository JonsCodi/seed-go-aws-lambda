package cmd

import (
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use: "start",

	Short: "This command will show you a hello world message",

	Long: `Welcome in start command, this cmd will display to you a hello world message.`,
}

func init() {
	rootCmd.AddCommand(startCmd)
}
