package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/JonsCodi/bava-go/internal/execute"
	"github.com/JonsCodi/bava-go/internal/model"
	"github.com/JonsCodi/bava-go/internal/model/packages"
	"github.com/JonsCodi/bava-go/pkg"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

// lambdaCmd represents the lambda command

const (
	apiGW = "APIGW"
	alb   = "ALB"
	event = "Event"
)

var lambdaCmd = &cobra.Command{
	Use: "lambda",
	Run: func(cmd *cobra.Command, args []string) {
		validateFlag(cmd.Flag("env-file").Value.String(), "flag --env-file or -e is required")
		validateFlag(cmd.Flag("name").Value.String(), "flag --name or -n is required")
		validateFlag(cmd.Flag("product").Value.String(), "flag --product or -p is required")
		validateFlag(cmd.Flag("kind").Value.String(), "flag --kind or -k is required")
		if err := isValidKind(cmd.Flag("kind").Value.String()); err != nil {
			println(err.Error())

			os.Exit(1)
		}

		envFlagValue, err := cmd.Flags().GetString("env-file")
		pkg.CheckErr(err)
		executeCommand := execute.New(bufio.NewReader(os.Stdin))

		//build folder and files
		name := cmd.Flag("name").Value.String()
		product := cmd.Flag("product").Value.String()
		description := cmd.Flag("description").Value.String()
		kind := cmd.Flag("kind").Value.String()
		project := model.New(name, product, description, kind, executeCommand.ExtractContent(envFlagValue))

		println("1. Creating structure folders...")
		//1. structure packages folder.
		if err := os.MkdirAll(fmt.Sprintf("./%s/internal", name), fs.ModePerm); err != nil {
			panic(err)
		}
		if err := os.Mkdir(fmt.Sprintf("./%s/internal/config", name), fs.ModePerm); err != nil {
			panic(err)
		}
		if err := os.Mkdir(fmt.Sprintf("./%s/internal/environment", name), fs.ModePerm); err != nil {
			panic(err)
		}
		println("1. Creating structure folders...Done")

		println("2. Creating go files...")
		//2. files of packages.
		mainFile, err := os.Create(fmt.Sprintf("./%s/main.go", name))
		pkg.CheckErr(err)
		defer mainFile.Close()
		envFile, err := os.Create(fmt.Sprintf("./%s/internal/environment/environment.go", name))
		pkg.CheckErr(err)
		defer envFile.Close()
		configFile, err := os.Create(fmt.Sprintf("./%s/internal/config/config.go", name))
		pkg.CheckErr(err)
		defer configFile.Close()
		modFile, err := os.Create(fmt.Sprintf("./%s/go.mod", name))
		pkg.CheckErr(err)
		defer modFile.Close()
		println("2. Creating go files...Done")

		println("3. Generating code...Done")
		pkg.CheckErr(packages.MainPackageTemplate.Execute(mainFile, project))
		pkg.CheckErr(packages.EnvPackageTemplate.ExecuteTemplate(envFile, "env.tmpl", project))
		pkg.CheckErr(packages.ConfigPackageTemplate.Execute(configFile, project))
		pkg.CheckErr(packages.GoModFileTemplate.Execute(modFile, project))
		println("3. Generating code...Done")
	},
}

func init() {
	startCmd.AddCommand(lambdaCmd)

	startCmd.PersistentFlags().StringP("kind", "k", "", "The aws event kind. Can be one of 'ALB' or 'APIGW' or 'Event' for a custom events")
}

func isValidKind(kind string) error {
	switch kind {
	case apiGW, alb, event:
		return nil
	default:
		return errors.New(fmt.Sprintf("%s is a invalid type, check -h flag for help", kind))
	}
}
