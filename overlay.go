package main

import (
	"log"
	"os"
	"text/template"

	"strings"
)

const (
	master  = `Names:{{define "list" }}{{"\n"}}{{range .}}{{println "-" .}}{{end}}{{end}}{{template "list" .}}`
	overlay = `{{define "list"}} {{join . ", "}}{{end}} `
)

var (
	funcs     = template.FuncMap{"join": strings.Join}
	guardians = []string{"Gamora", "Groot", "Nebula", "Rocket", "Star-Lord"}
)

func main() {
	masterTmpl := template.Must(template.New("master").Funcs(funcs).Parse(master))
	overlayTmpl, err := template.Must(masterTmpl.Clone()).Parse(overlay)

	if err != nil {
		log.Fatal(err)

	}
	if err := masterTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
	if err := overlayTmpl.Execute(os.Stdout, guardians); err != nil {
		log.Fatal(err)
	}
}
