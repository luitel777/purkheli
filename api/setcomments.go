package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/luitel777/purkheli/utils"
)

// url: /api/setcomments
func (V Values) setComments(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 40)
	if err != nil {
		http.Error(w, "Error while parsing\r\n"+err.Error(), http.StatusBadRequest)
		return
	}

	title := r.Form.Get("title")
	textarea := r.Form.Get("textarea")
	id := r.Form.Get("postid")

	err = utils.ValidateTitle(title)
	if err != nil {
		http.Error(w, "Error on validating title", http.StatusBadRequest)
		return
	}

	err = utils.ValidateTextArea(textarea)
	if err != nil {
		http.Error(w, "Error on validating textarea", http.StatusBadRequest)
		return
	}
	id_int, err := strconv.Atoi(id)
	if err != nil {
		log.Println("line 38: api/setcomments.go cannot convert " + id + " Atoi")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		V.db.StoreComments(title, textarea, id_int, "nil")
	} else {
		filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), header.Filename)

		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)

		V.db.StoreComments(title, textarea, id_int, filename[8:])

		f.Close()
		file.Close()
	}

	w.Write([]byte("successfully submitted the comment on id " + id))
	V.db.CloseDB()
	return
}
