package controladores

import (
	"encoding/json"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"
)

func EventosGet(w http.ResponseWriter, r *http.Request) {
	var eventos []modelos.Evento
	err := db.Database.Find(&eventos)
	if err != nil {
		http.Error(w, "Error al hacer la consulta al servidor", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := json.NewEncoder(w).Encode(eventos)
	if response != nil {
		http.Error(w, "Error al hacer la codificacion", http.StatusInternalServerError)
		return
	}
}
