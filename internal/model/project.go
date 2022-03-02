package model

import (
	"fmt"
	"io/fs"
	"os"
)

type Project struct {
	Name        string
	Product     string
	Description string
}

func New(name, product string) Project {
	return Project{
		Name:    name,
		Product: product,
	}
}

func (p Project) MakeModule() (err error) {
	mod := []byte(fmt.Sprintf(moduleContent, p.Name))
	modFile, err := os.Create(fmt.Sprintf("./%s/go.mod", p.Name))
	_, err = modFile.Write(mod)
	defer modFile.Close()

	return
}

func (p Project) MakeMain() (err error) {
	err = os.Mkdir(p.Name, fs.ModePerm)
	if err != nil {
		return
	}

	mainGo := []byte(fmt.Sprintf(mainContentForLambda, p.Name, p.Product))
	mainGoFile, err := os.Create(fmt.Sprintf("./%s/main.go", p.Name))
	if err != nil {
		return
	}
	_, err = mainGoFile.Write(mainGo)
	defer mainGoFile.Close()

	return
}

func (p Project) MakeHandler(event, res string) (err error) {
	err = os.Mkdir(fmt.Sprintf("./%s/internal", p.Name), fs.ModePerm)
	if err != nil {
		return
	}

	if res == "" {
		err = os.Mkdir(fmt.Sprintf("./%s/internal/model", p.Name), fs.ModePerm)
		eventModel := []byte(awsEventmodel)
		eventModelGoFile, err := os.Create(fmt.Sprintf("./%s/internal/model/event.go", p.Name))
		if err != nil {
			return err
		}

		_, err = eventModelGoFile.Write(eventModel)
		if err != nil {
			return err
		}
		defer eventModelGoFile.Close()
	}
	handlerGo := []byte(fmt.Sprintf(handlerContentLambda, event, res, p.Name, p.Product))
	handlerGoFile, err := os.Create(fmt.Sprintf("./%s/internal/handler.go", p.Name))
	_, err = handlerGoFile.Write(handlerGo)
	defer handlerGoFile.Close()

	return
}
