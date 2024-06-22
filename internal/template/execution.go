package template

import (
	"bytes"
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, tmpl *template.Template, data any) {
	var buf bytes.Buffer
	err := tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	buf.WriteTo(w)

}
