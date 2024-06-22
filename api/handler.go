package api

import (
	"net/http"
	"strings"

	"github.com/luitel777/purkheli/utils"
)

type ApiHandler struct {
	v Values
}

// serve everything from /api in here
func (A *ApiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// initating this handler also initates the database
	// the database pool is never closed unless the server
	// is shut down
	A.v.db.InitiateDB()

	// remove ending trails since /api/submit or /api/submit/ are same here
	path := utils.RemoveTrail(r.URL.Path)

	switch path {
	case "/api/":
		data := []byte("please visit docs to use this API")
		w.Write(data)
	case "/api/getposts":
		A.v.getPosts(w, r)
	case "/api/submit":
		A.v.submit(w, r)
	case "/api/retrive_post":
		A.v.getSinglePost(w, r)
	case "/api/delete":
		A.v.delete(w, r)
	default:
		// I know this code formatting looks abysmal but > _ < I don't know
		// best practices
		if strings.HasPrefix(path, "/api/getcomments") {
			A.v.getComments(w, r)
		} else if strings.HasPrefix(path, "/api/setcomments") {
			// w.Write([]byte("sss"))
			A.v.setComments(w, r)
		} else {
			w.Write([]byte("error"))
		}
	}
}
