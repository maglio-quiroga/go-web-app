package main

import (
	"go-web-app/controladores"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	db.Iniciar_conexion()
	db.Database.AutoMigrate(modelos.Alumno{})
	db.Database.AutoMigrate(modelos.Evento{})
	db.Database.AutoMigrate(modelos.Empleo{})
	db.Database.AutoMigrate(modelos.Admin{})
	db.Database.AutoMigrate(modelos.Peticion{})

	enrutador := mux.NewRouter()
	enrutador.HandleFunc("/log", controladores.LogAdmin).Methods("POST")
	http.ListenAndServe(":3000", enrutador)
}
