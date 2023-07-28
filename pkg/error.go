package pkg

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorBody struct {
	Message string
	Code    int
}

func NotFound(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNotFound)

	error := ErrorBody{
		Message: "Not Found",
		Code:    404,
	}

	err = tmp.Execute(w, error)
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}
}

func MethodNotAllowed(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error.html")
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)

	error := ErrorBody{
		Message: "Method Not Allowed",
		Code:    405,
	}

	err = tmp.Execute(w, error)
	if err != nil {
		log.Println(err.Error())
		InternalServerError(w)
		return
	}
}

func InternalServerError(w http.ResponseWriter) {
	tmp, err := template.ParseFiles("./ui/html/error500.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)

	err = tmp.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
