package packages

import "text/template"

var MainPackageTemplate = template.Must(template.New("main.go").
	Parse(`package main

func main() {
	println("Hello world by bava-go. Welcome to {{ .Name }} of {{ .Product }}")	
}
`))
