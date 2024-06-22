package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/luitel777/purkheli/internal/db"
	"github.com/luitel777/purkheli/utils"
)

type Values struct {
	db db.Database
}

func (V Values) submit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 40)
	if err != nil {
		http.Error(w, "Error while parsing\r\n"+err.Error(), http.StatusBadRequest)
		return
	}
	title := r.Form.Get("title")
	textarea := r.Form.Get("textarea")

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

	file, header, err := r.FormFile("file")
	if err != nil {
		V.db.StoreDB(title, textarea, "nil")
	} else {
		filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), header.Filename)

		f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		io.Copy(f, file)

		V.db.StoreDB(title, textarea, filename[8:])

		f.Close()
		file.Close()
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("successfully submitted the post"))
	V.db.CloseDB()
	return
}
