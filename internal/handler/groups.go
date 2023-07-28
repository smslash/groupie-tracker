package handler

import (
	"log"
	"net/http"
	"text/template"

	"git/ssengerb/groupie-tracker/pkg"
)

func Groups(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/groups/" {
		pkg.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/groups.html")
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	err = tmp.Execute(w, pkg.Bands)
	if err != nil {
		log.Println("Error:", err)
		pkg.InternalServerError(w)
		return
	}
}
