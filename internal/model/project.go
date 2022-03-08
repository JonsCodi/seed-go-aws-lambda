package model

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"io/fs"
	"os"
	"strings"
)

type EnvironmentField struct {
	Attr string
	Env  string
}

type Project struct {
	Name              string
	Product           string
	Description       string
	EnvironmentFields []EnvironmentField
	kind              string
}

func New(name, product, description, kind, envContentFile string) Project {
	environmentFields := strings.Replace(envContentFile, "=''", "", -1)
	fields := strings.Split(environmentFields, "\n")

	envFields := make([]EnvironmentField, len(fields))
	for i, e := range fields {
		envFields[i] = EnvironmentField{
			Attr: strcase.ToCamel(strings.ToLower(e)),
			Env:  e,
		}
	}

	return Project{
		Name:              name,
		Product:           product,
		Description:       description,
		EnvironmentFields: envFields,
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
