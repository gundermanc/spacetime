// List of templates to embed in the client app as strings.

package server

//go:generate go get github.com/wlbr/templify
//go:generate mkdir ../generated
//go:generate go run github.com/wlbr/templify -p generated -o ../generated/index.generated.go -e index.gohtml
