package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func main() {
	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.HandleFunc("/", Index)
	http.ListenAndServe(":3000", enrutador)
}
