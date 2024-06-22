package template

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)

// url: /create
func CreateNewPost(w http.ResponseWriter, r *http.Request) {

	fmt.Println(w.Header().Get("error"))

	tmpl, err := template.ParseFiles("web/create.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Println("Error at internal/template/create.go", err)
		return
	}
	ExecuteTemplate(w, tmpl, nil)
}
