package packages

import (
	"embed"
	"github.com/JonsCodi/bava-go/pkg"
	"text/template"
)

//go:embed *.tmpl
var tpls embed.FS
var err error
var EnvPackageTemplate *template.Template

func init() {
	EnvPackageTemplate, err = template.ParseFS(tpls, "*")
	pkg.CheckErr(err)
}
