//  Server daemon

package server

import (
	"net/http"
)

// StartRemoteAccessServer starts a remote access server that
// listens for connections to this machine over HTTP.
func StartRemoteAccessServer() error {

	// Setup file server for HTTP root.
	registerAssetFileServers()

	// Setup specific page handlers.
	registerIndexHandler()

	return http.ListenAndServe(":8080", nil)
}
