package handler

import (
	"log"
	"net/http"
	"text/template"

	"git/ssengerb/groupie-tracker/pkg"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search/" {
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

	value := r.FormValue("query")
	list := pkg.SearchFor(value)

	err = tmp.Execute(w, list)
	if err != nil {
		log.Println("Error:", err)
		pkg.InternalServerError(w)
		return
	}
}
