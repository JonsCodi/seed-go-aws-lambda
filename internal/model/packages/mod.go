package packages

import "text/template"

var GoModFileTemplate = template.Must(template.New("go.mod").Parse(
	`module {{ .Name}}

go 1.17`))
