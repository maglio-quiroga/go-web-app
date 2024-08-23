package main

import (
	"go-web-app/db"
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

	db.Iniciar_conexion()

	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.HandleFunc("/", Index)
	http.ListenAndServe(":3000", enrutador)
}
