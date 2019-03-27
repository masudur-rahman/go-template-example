package main

import (
	"html/template"
	"log"
	"os"
)

type Data struct {
	True  bool
	False bool
}

const tmplBuiltin = `Or: {{ if or .True .False }}true{{ else }}false{{ end }}
And: {{ if and .True .False }}true{{ else}}false{{ end }}
Not: {{ if not .False }}true{{ else }}false{{ end }}
`

const tmplIndex = `Slice[0]: {{ index .Slice 0 }}
SliceNested[1][0]: {{ index .SliceNested 1 0 }}
Map["key"]: {{ index .Map "key" }}`

func main() {
	t := template.Must(template.New("and_or_not").Parse(tmplBuiltin))
	data := Data{True: true, False: false}

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

	tIndex := template.Must(template.New("index").Parse(tmplIndex))
	data2 := struct {
		Slice       []string
		SliceNested [][]int
		Map         map[string]int
	}{
		Slice: []string{"first", "second"},
		SliceNested: [][]int{
			{3, 1},
			{2, 1},
		},
		Map: map[string]int{
			"key": 5,
		},
	}
	if err := tIndex.Execute(os.Stdout, data2); err != nil {
		log.Fatal(err)
	}
}
