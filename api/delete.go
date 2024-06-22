package api

import (
	"net/http"
)

// to delete the post you need to have a passcode in the passcodes table
// this should be done manually with database
func (V Values) delete(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil{
		http.Error(w, "Error while parsing\r\n"+err.Error(), http.StatusBadRequest)
		return
	}
	postid := r.Form.Get("id")
	passcode := r.Form.Get("passcode")
	err = V.db.DeletePost(postid, passcode)

	if err != nil{
		http.Error(w, "Cannot delete the database", http.StatusMethodNotAllowed)
	}
}
