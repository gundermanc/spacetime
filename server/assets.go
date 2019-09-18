package server

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
)

func registerAssetFileServers() {

	// Map static assets path.
	httpBox := rice.MustFindBox("http")
	httpFs := http.StripPrefix("/static/", http.FileServer(httpBox.HTTPBox()))
	http.Handle("/static/", httpFs)

	// Map bootstrap path.
	bootstrapBox := rice.MustFindBox("../node_modules/bootstrap/dist")
	bootstrapFs := http.StripPrefix("/static/bootstrap/", http.FileServer(bootstrapBox.HTTPBox()))
	http.Handle("/static/bootstrap/", bootstrapFs)
}
