package packages

import "text/template"

var EnvPackageTemplate = template.Must(template.New("environment.go").
	Parse(`package environment

type Environment struct {
{{- range .EnvironmentFields }}
	{{ . }} string env:"{{ . }}"
{{- end }}
}
`))
