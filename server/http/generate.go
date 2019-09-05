// List of templates to embed in the client app as strings.

package http

// Templates: -------
//go:generate go get github.com/wlbr/templify
//go:generate mkdir ../generated
//go:generate go run github.com/wlbr/templify -p generated -o ../generated/index.generated.go -e index.gohtml

// Assets: -------
//go:generate go run github.com/wlbr/templify -p generated -o ../generated/bootstrap.generated.go -e ../../node_modules/bootstrap/dist/css/bootstrap.min.css
//go:generate go run github.com/wlbr/templify -p generated -o ../generated/bootstrapmap.generated.go -e ../../node_modules/bootstrap/dist/css/bootstrap.min.css.map
