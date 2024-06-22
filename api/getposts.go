package api

import (
	"encoding/json"
	"net/http"
	"sort"
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

func (V Values) getPosts(w http.ResponseWriter, _ *http.Request) {
	err, items := V.db.RetriveDb()
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

		// append is not always deterministic
		// sometime it appends 2nd or 3rd value first and
		// remaining later which makes the return value inconsistent
		// therfore we sort as soon as we get the data
		// if true i should come before j
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
