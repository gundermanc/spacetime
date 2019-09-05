// Serving utilites

package utils

import (
	"html/template"
	"io"
	"log"
)

func WriteTemplate(w io.Writer, tmplName string, tmplBody string, model interface{}) {
	tmpl, err := template.New(tmplName).Parse(tmplBody)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, model)
	if err != nil {
		log.Fatal(err)
	}
}
