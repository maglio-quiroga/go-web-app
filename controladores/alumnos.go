package controladores

import (
	"encoding/json"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"
	"strconv"
)

func AlumnosGet(w http.ResponseWriter, r *http.Request) {
	PageNumStr := r.URL.Query().Get("pagina")
	if PageNumStr == "" {
		PageNumStr = "1"
	}
	PageNum, err := strconv.Atoi(PageNumStr)
	if err != nil || PageNum < 1 {
		PageNum = 1
	}
	const MaxEmpleos = 4
	alternador := (PageNum - 1) * MaxEmpleos
	var empleos []modelos.Empleo
	resp := db.Database.Limit(MaxEmpleos).Offset(alternador).Find(&empleos)
	if resp.Error != nil {
		http.Error(w, "Error al obtener los empleos", http.StatusInternalServerError)
		return
	}
	var EmpleoCount int64
	db.Database.Model(&modelos.Empleo{}).Count(&EmpleoCount)
	TotalEmpleos := (int(EmpleoCount) + MaxEmpleos - 1) / MaxEmpleos
	estructura := map[string]interface{}{
		"empleos":      empleos,
		"totalPaginas": TotalEmpleos,
		"paginaActual": PageNum,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(estructura); err != nil {
		http.Error(w, "Error al codificar la respuesta", http.StatusInternalServerError)
	}
}
