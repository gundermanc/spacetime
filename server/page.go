// Base page infrastructure

package server

type pageViewModel struct {
	Title string
}

func newPageViewModel(title string) pageViewModel {
	return pageViewModel{title}
}
