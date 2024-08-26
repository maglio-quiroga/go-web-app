package controladores

import (
	"encoding/json"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"
)

func EmpleosGet(w http.ResponseWriter, r *http.Request) {
	var empleos []modelos.Empleo
	err := db.Database.Find(&empleos).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := json.NewEncoder(w).Encode(empleos)
	if response != nil {
		http.Error(w, response.Error(), http.StatusInternalServerError)
		return
	}
}
