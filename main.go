package main

import (
	"go-web-app/controladores"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	db.Iniciar_conexion()
	db.Database.AutoMigrate(modelos.Alumno{})
	db.Database.AutoMigrate(modelos.Evento{})
	db.Database.AutoMigrate(modelos.Empleo{})
	db.Database.AutoMigrate(modelos.Admin{})
	db.Database.AutoMigrate(modelos.Peticion{})

	enrutador := mux.NewRouter()
	enrutador.Use(corsMiddleware)
	enrutador.PathPrefix("./aplicacion/").Handler(http.StripPrefix("/aplicacion/", http.FileServer(http.Dir("./aplicacion/"))))

	enrutador.HandleFunc("/inicio", controladores.InicioGet).Methods("GET")
	enrutador.HandleFunc("/inicio", controladores.InicioPost).Methods("POST")
	enrutador.HandleFunc("/empleos", controladores.EmpleosGet).Methods("GET")
	enrutador.HandleFunc("/eventos", controladores.EventosGet).Methods("GET")

	enrutador.HandleFunc("/log", controladores.LogAdmin).Methods("POST")
	enrutador.HandleFunc("/alumnos", controladores.AlumnosGet).Methods("GET")

	http.ListenAndServe(":8080", enrutador)
}
