package main

import (
	"go-web-app/controladores"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"

	"github.com/gorilla/mux"
)

func ManejadorRutas(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		//enviar json para el index.html
	case "/alumnos":
		//enviar json para la pagina del listado de alumnos
	case "/empleos":
		//enviar json para la pagina del listado de empleos
	case "/adm-peticiones":
		//enviar json para la pagina de administrador de las peticiones
	case "/adm-alumnos":
		//
	case "/adm-eventos":
		//
	case "/adm-empleos":
		//
	case "/form-log":
		controladores.LogAdmin(w, r)
	}
	//Los Json seran recuperados en el cliente con fetch + async await.
}

func main() {

	db.Iniciar_conexion()
	db.Database.AutoMigrate(modelos.Alumno{})
	db.Database.AutoMigrate(modelos.Evento{})
	db.Database.AutoMigrate(modelos.Empleo{})
	db.Database.AutoMigrate(modelos.Admin{})
	db.Database.AutoMigrate(modelos.Peticion{})

	enrutador := mux.NewRouter()
	enrutador.HandleFunc("/", ManejadorRutas)
	http.ListenAndServe(":3000", enrutador)
}
