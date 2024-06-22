package template

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/luitel777/purkheli/utils"
)

type CommentInformation struct {
	Comments  []Data
	DataField Param
	PostId    int
}

func Comments(w http.ResponseWriter, r *http.Request) {
	commentFields := CommentInformation{}
	path := utils.RemoveTrail(r.URL.Path)

	if len(path) < 7 {
		w.Write([]byte("illegal path"))
		http.Error(w, "Error: wrong ID", http.StatusBadRequest)
		return
	} else {
		id, err := strconv.Atoi(path[7:])

		url := "http://localhost:9000/api/retrive_post?id=%d"
		url = fmt.Sprintf(url, id)
		err, commentFields.DataField = ConvertJsonSingle(url)
		if err != nil {
			http.Error(w, "Error converting json", http.StatusBadRequest)
			return
		}

		url = "http://localhost:9000/api/getcomments?id=%d"
		url = fmt.Sprintf(url, id)
		err, commentFields.Comments = ConvertJson(url)
		if err != nil {
			http.Error(w, "Error converting json", http.StatusBadRequest)
			return
		}

		commentFields.PostId = id

		tmpl, err := template.ParseFiles("web/comments.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			log.Println("Error at internal/template/comments.go", err)
			return
		}
		ExecuteTemplate(w, tmpl, commentFields)
	}

}

