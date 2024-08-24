package main

import (
	"go-web-app/db"
	"go-web-app/modelos"
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
	db.Database.AutoMigrate(modelos.Alumno{})
	db.Database.AutoMigrate(modelos.Evento{})
	db.Database.AutoMigrate(modelos.Empleo{})
	db.Database.AutoMigrate(modelos.Admin{})
	db.Database.AutoMigrate(modelos.Peticion{})

	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.HandleFunc("/", Index)
	http.ListenAndServe(":3000", enrutador)
}
