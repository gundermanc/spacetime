package server

import (
	"net/http"

	"github.com/gundermanc/spacetime/utils"

	"github.com/gundermanc/spacetime/server/generated"
)

func registerAssetRoutes() {
	http.HandleFunc("/bootstrap.min.css", bootstrapHandler)
	http.HandleFunc("/bootstrap.min.css.map", bootstrapMapHandler)
}

func bootstrapHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	utils.Write(w, generated.BootstrapTemplate())
}

func bootstrapMapHandler(w http.ResponseWriter, r *http.Request) {
	utils.Write(w, generated.BootstrapmapTemplate())
}
