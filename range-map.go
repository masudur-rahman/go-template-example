package main

import (
	"html/template"
	"log"
	"os"
)

const tmplRange = `Elements of map:
{{ range $k, $v := . }}{{ $k }}: {{ $v }}
{{end}}`

const tmplChan = `Elements of a channel: {{ range . }}{{ . }} {{end}}`

func main() {
	t := template.Must(template.New("range").Parse(tmplRange))

	data := map[string]int{
		"one":  1,
		"five": 5,
	}
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}

	tt := template.Must(template.New("channel").Parse(tmplRange))
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()
	if err := tt.Execute(os.Stdout, ch); err != nil {
		log.Fatal(err)
	}
}
