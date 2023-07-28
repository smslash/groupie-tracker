package handler

import (
	"log"
	"net/http"
	"text/template"

	"git/ssengerb/groupie-tracker/pkg"
)

func Filters(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/filters/" {
		pkg.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/search.html")
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	list := pkg.FilterFor(r)

	err = tmp.Execute(w, list)
	if err != nil {
		log.Println("Error:", err)
		pkg.InternalServerError(w)
		return
	}
}
