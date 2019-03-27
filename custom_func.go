package main

import (
	"log"
	"os"
	"text/template"
)

func sum(x, y int) int {
	return x + y
}

const tmplFunc = `5 + {{ .Arg }} = {{ sum 5 .Arg }}`

func main() {
	customFunctions := template.FuncMap{
		"sum": sum,
	}
	tFunc := template.Must(template.New("func").Funcs(customFunctions).Parse(tmplFunc))

	data := struct {
		Arg int
	}{
		Arg: 6,
	}
	if err := tFunc.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
