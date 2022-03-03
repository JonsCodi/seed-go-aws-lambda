package execute

import (
	"bufio"
	"github.com/JonsCodi/bava-go/pkg"
	"io/ioutil"
	"strings"
)

type CommandInterface interface {
	AskForInput(question string) string
	ExtractContent(fileName string) string
}

type command struct {
	reader *bufio.Reader
}

func New(reader *bufio.Reader) CommandInterface {
	return &command{reader}
}

func (c command) AskForInput(question string) string {
	print(question)
	input, err := c.reader.ReadString('\n')
	pkg.CheckErr(err)

	return strings.Replace(input, "\n", "", 1)
}

func (c command) ExtractContent(fileName string) string {
	out, err := ioutil.ReadFile(fileName)
	pkg.CheckErr(err)

	return string(out)
}
