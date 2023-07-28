package handler

import (
	"log"
	"net/http"
	"text/template"

	"git/ssengerb/groupie-tracker/pkg"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		pkg.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/home.html")
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println("Error:", err)
		pkg.InternalServerError(w)
		return
	}
}
