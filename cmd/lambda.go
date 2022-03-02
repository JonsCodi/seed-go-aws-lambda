package cmd

import (
	"bufio"
	"fmt"
	"github.com/JonsCodi/bava-go/internal/model"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use: "lambda",
	Run: func(cmd *cobra.Command, args []string) {
		//build folder and files
		product := model.New(getNameAndProduct())
		err := product.MakeMain()
		check(err)
		err = product.MakeModule()
		check(err)
		err = product.MakeHandler(getEvent())
		check(err)
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

func getEvent() (event, res string) {
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
