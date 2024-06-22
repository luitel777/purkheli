package server

import (
	"net/http"

	"github.com/luitel777/purkheli/internal/template"
)

type PurkheliHandler struct {
}

// maps / path to a html template page
func (P *PurkheliHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		template.Homepage(w)
	case "/create":
		template.CreateNewPost(w, r)
	case "/index":
		data := []byte("index")
		w.Write(data)
	default:
		data := []byte("error")
		w.Write(data)
	}
}
