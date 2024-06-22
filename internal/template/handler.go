package template

import (
	"net/http"

	"github.com/luitel777/purkheli/utils"
)

type TemplateHandler struct{}

func (T *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Path
	url = utils.RemoveTrail(url)

	if url[len(url)-6:] == "create" {
		CreateNewComment(w, r)
	} else {
		Comments(w, r)
	}
}
