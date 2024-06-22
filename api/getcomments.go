package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func (V Values) getComments(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	post_id, err := strconv.Atoi(id)
	if err != nil {
		log.Println("line 14 : api/getcomments.go cannot convert " + id + " Atoi")
	}

	err, items := V.db.RetriveCommentsDb(post_id)
	var dataList []Data
	for i, j := range items {
		data := Data{
			Id: i,
			Params: Param{
				Title:    j[0],
				Textarea: j[1],
				Imgpath:  j[2],
				Time:     j[3],
			},
		}

		dataList = append(dataList, data)

		sort.Slice(dataList, func(i, j int) bool {
			return dataList[i].Id > dataList[j].Id
		})

	}

	responseData, err := json.Marshal(dataList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseData)
}
