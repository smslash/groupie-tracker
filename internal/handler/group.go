package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"git/ssengerb/groupie-tracker/pkg"
)

func Group(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/group/" {
		pkg.NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		pkg.MethodNotAllowed(w)
		return
	}

	tmp, err := template.ParseFiles("./ui/html/group.html")
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 || id > 52 {
		pkg.NotFound(w)
		return
	}

	pkg.Bands[id-1].Coordinate = pkg.Geolocalization(pkg.Bands[id-1])

	json, err := json.Marshal(pkg.Bands[id-1].Coordinate)
	if err != nil {
		pkg.InternalServerError(w)
		return
	}

	pkg.Bands[id-1].EncodeLoc = string(json)
	err = tmp.Execute(w, pkg.Bands[id-1])
	if err != nil {
		log.Println("Error:", err)
		pkg.InternalServerError(w)
		return
	}
}
