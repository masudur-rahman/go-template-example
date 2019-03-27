package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Okay struct {
	Name   string
	IsOkay bool
	Reason string
}

const tmplStr = `{{range . -}}
{{if .IsOkay}}'{{.Name}}' is Okay, Reason: '{{.Reason}}'.{{else}}'{{.Name}} is not Okay, Reason: '{{.Reason}}.'{{end}}
{{end}}`

const tmplStr2 = `{{- define "okay" -}}'{{.Name}}' is Okay, Reason: '{{.Reason}}'.{{end}}
{{- define "notOkay" -}}'{{.Name}} is not Okay, Reason: '{{.Reason}}'.{{end}}
{{- define "checkOkay" -}}{{if .IsOkay}}{{template "okay" .}}{{else}}{{ template "notOkay" . }}{{end}}{{end}}

{{- range . -}}
{{template "checkOkay" .}}
{{end}}`

func main() {
	tmpl := template.Must(template.New("okay").Parse(tmplStr))
	tmpl2 := template.Must(template.New("okay2").Parse(tmplStr2))

	data := []Okay{
		Okay{
			"Masudur Rahman",
			true,
			"He is happy",
		},
		Okay{
			"Kamol Hasan",
			false,
			"He is unhappy",
		},
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("< ========================================= >")
	if err := tmpl2.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

}
