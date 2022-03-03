package cmd

import (
	"bufio"
	"fmt"
	"github.com/JonsCodi/bava-go/internal/execute"
	"github.com/JonsCodi/bava-go/internal/model"
	"github.com/JonsCodi/bava-go/internal/model/packages"
	"github.com/JonsCodi/bava-go/pkg"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"strings"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use: "lambda",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("env-file").Value.String() == "" {
			print("flag --enf-file or -e is required")

			os.Exit(1)
		}
		envFlagValue, err := cmd.Flags().GetString("env-file")
		pkg.CheckErr(err)
		executeCommand := execute.New(bufio.NewReader(os.Stdin))

		envFileContent := executeCommand.ExtractContent(envFlagValue)
		fields := strings.Split(envFileContent, "=")
		//build folder and files
		name := executeCommand.AskForInput("Project name? ")
		product := executeCommand.AskForInput("Product name? ")
		description := executeCommand.AskForInput("Any description about it? ")
		project := model.New(name, product, description, fields)

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

		check(packages.MainPackageTemplate.Execute(mainFile, project))
		check(packages.EnvPackageTemplate.Execute(envFile, project))
		check(packages.ConfigPackageTemplate.Execute(configFile, project))
	},
}

var mapOfAWSProjects map[string]string

func init() {
	startCmd.AddCommand(lambdaCmd)
	mapOfAWSProjects = map[string]string{
		"A": "events.APIGatewayProxyRequest",
		"B": "events.ALBTargetGroupRequest",
		"C": "model.AwsEvent",
	}
}

func getNameAndProduct() (name, product string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Name: ")
	name, err := reader.ReadString('\n')
	check(err)

	fmt.Print("Product: ")
	product, err = reader.ReadString('\n')
	check(err)

	name = strings.Replace(name, "\n", "", -1)
	product = strings.ToLower(strings.Replace(product, "\n", "", -1))

	return
}

func getEvent(option string) (event, res string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What Kind of event? Press a option to choose")
	for option, event := range mapOfAWSProjects {
		fmt.Println(option + " - " + event)
	}

	option, err := reader.ReadString('\n')
	check(err)
	option = strings.Replace(option, "\n", "", -1)
	switch option {
	case "A":
		return mapOfAWSProjects["A"], "(*events.APIGatewayProxyResponse, error)"
	case "B":
		return mapOfAWSProjects["B"], "(*events.ALBTargetGroupResponse, error)"
	case "C":
		return mapOfAWSProjects["C"], ""
	default:
		panic("invalid option")
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
