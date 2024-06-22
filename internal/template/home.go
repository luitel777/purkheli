package template

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type Param struct {
	Title    string `json:"Title"`
	Textarea string `json:"Textarea"`
	Imgpath  string `json:"Imgpath"`
	Time     string `json:"Time"`
}

type Data struct {
	Id     int
	Params Param
}

// we get all the data from /api/getposts endpoint
func GetPostFromEndpoint(endpoint string) (error, []byte) {
	res, err := http.Get(endpoint)
	if err != nil {
		return err, []byte("failed to get data")
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err, []byte("failed to read data")
	}
	return err, data
}

func ConvertJsonSingle(endpoint string) (error, Param) {
	d := Param{}
	err, p := GetPostFromEndpoint(endpoint)
	if err != nil {
		return err, d
	}
	err = json.Unmarshal(p, &d)
	return err, d
}

func ConvertJson(endpoint string) (error, []Data) {
	d := []Data{}
	err, p := GetPostFromEndpoint(endpoint)
	if err != nil {
		return err, d
	}
	err = json.Unmarshal(p, &d)
	return err, d
}

// url :/
func Homepage(w http.ResponseWriter) {
	err, fields := ConvertJson("http://localhost:9000/api/getposts")
	if err != nil {
		http.Error(w, "Cannot evaluate posts for some reason: " + err.Error(), http.StatusInternalServerError)
		fields = []Data{}
		return
	}
	tmpl, err := template.ParseFiles("web/index.html")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		log.Println("Error at internal/template/home.go", err)
		return
	}
	ExecuteTemplate(w, tmpl, fields)
}
