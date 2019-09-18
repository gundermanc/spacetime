// Index page code behind

package server

import (
	"net/http"
)

type indexViewModel struct {
	pageViewModel

	Header string
}

func newIndexViewModel(title string, header string) indexViewModel {
	base := newPageViewModel(title)

	return indexViewModel{base, header}
}

func registerIndexHandler() {
	http.HandleFunc("/", serverRoot)
}

func serverRoot(w http.ResponseWriter, r *http.Request) {
	writeTemplate(w, "index.gohtml", newIndexViewModel(appName, indexHeader))
}
