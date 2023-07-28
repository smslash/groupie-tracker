package internal

import (
	"log"
	"net/http"
	"time"

	"git/ssengerb/groupie-tracker/internal/handler"
	"git/ssengerb/groupie-tracker/pkg"
)

func Server() {
	mux := http.NewServeMux()
	pkg.Parse()
	pkg.FilterConfiguration()

	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/groups/", handler.Groups)
	mux.HandleFunc("/group/", handler.Group)
	mux.HandleFunc("/search/", handler.Search)
	mux.HandleFunc("/filters/", handler.Filters)

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./ui/static/"))))

	server := http.Server{
		Addr:         ":4000",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("The server is running on http://127.0.0.1" + server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
