package api

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type singleData struct {
	Title    string `json:"Title"`
	Textarea string `json:"Textarea"`
	Imgpath  string `json:"Imgpath"`
	Time     string `json:"Time"`
}

func (V Values) getSinglePost(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	err, items := V.db.RetriveSinglePost(id)

	data := singleData{}
	data.Title = items[0]
	data.Textarea = items[1]
	data.Imgpath = items[2]
	data.Time = items[3]

	responseData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
