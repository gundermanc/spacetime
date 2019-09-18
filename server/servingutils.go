// Serving utilites

package server

import (
	"html/template"
	"io"
	"log"

	rice "github.com/GeertJohan/go.rice"
)

func writeTemplate(w io.Writer, tmplPath string, model interface{}) {

	box, err := rice.FindBox("http")
	if err != nil {
		log.Fatal(err)
	}

	tmplStr, err := box.String(tmplPath)
	if err != nil {
		log.Fatal(err)
	}

	tmpl, err := template.New(tmplPath).Parse(tmplStr)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, model)
	if err != nil {
		log.Fatal(err)
	}
}

func write(w io.Writer, content string) {
	_, err := w.Write([]byte(content))
	if err != nil {
		log.Fatal(err)
	}
}
