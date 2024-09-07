package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	api "github.com/luitel777/purkheli/api"
	"github.com/luitel777/purkheli/internal/db"
	server "github.com/luitel777/purkheli/internal/server"
	"github.com/luitel777/purkheli/internal/template"
)

func main() {
	mux := http.NewServeMux()

	os.Mkdir("uploads", 755)

	db := db.Database{}
	db.InitiateDB()

	handler := &server.PurkheliHandler{}
	apihandler := &api.ApiHandler{}
	templatehandler := &template.TemplateHandler{}

	mux.Handle("/", handler)
	mux.Handle("/api/", apihandler)
	mux.Handle("/posts/", templatehandler)
	mux.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	serv := &http.Server{
		Addr:    "localhost:9000",
		Handler: mux,
	}

	go func() {
		fmt.Println("running server on http://localhost:9000")
		err := serv.ListenAndServe()
		if err != nil {
			log.Println(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("gracefully shutting down", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Println(err)
	}
	serv.Shutdown(tc)
}
