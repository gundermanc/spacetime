// Index page code behind

package server

import (
	"net/http"

	"github.com/gundermanc/spacetime/server/generated"
	"github.com/gundermanc/spacetime/utils"
)

//go:generate go run github.com/wlbr/templify -p server -o indextemp.go index.gohtml

type indexViewModel struct {
	pageViewModel

	Header string
}

func newIndexViewModel(title string, header string) indexViewModel {
	base := newPageViewModel(title)

	return indexViewModel{base, header}
}

func registerIndex() {
	http.HandleFunc("/", serverRoot)
}

func serverRoot(w http.ResponseWriter, r *http.Request) {

	utils.WriteTemplate(w, "index", generated.IndexTemplate(), newIndexViewModel(appName, indexHeader))
}
