package template

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/luitel777/purkheli/utils"
)

// url: /posts/id/create
func CreateNewComment(w http.ResponseWriter, r *http.Request) {
	var id int
	url := r.URL.Path
	url = utils.RemoveTrail(url)
	urls := strings.Split(url, "/")
	if urls[1] == "posts" && urls[3] == "create" {
		id, _ = strconv.Atoi(urls[2])
	}

	data := struct {
		Id int
	}{
		Id: id,
	}

	tmpl, err := template.ParseFiles("web/createComment.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Println("Error at internal/template/createComment.go", err)
		return
	}

	ExecuteTemplate(w, tmpl, data)
}
