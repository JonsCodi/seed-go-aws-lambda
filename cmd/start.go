package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.PersistentFlags().StringP("env-file", "e", "", "Config environment file")
	startCmd.PersistentFlags().StringP("name", "n", "", "Name of the project, example (gitlab.com/lambda_project")
	startCmd.PersistentFlags().StringP("product", "p", "", "Product name that project will be represented")
	startCmd.PersistentFlags().StringP("description", "d", "", "Any description about it")
}

func validateFlag(value, message string) {
	if value == "" {
		println(message)

		os.Exit(1)
	}
}
