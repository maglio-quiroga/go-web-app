package controladores

import (
	"encoding/json"
	"go-web-app/modelos"
	"log"
	"net/http"
)

func LogAdmin(w http.ResponseWriter, r *http.Request) {
	usuario := ObtenerUsuario(r)
	modelos.DefaultAdminServ.AutenticarAdmin(w, usuario)
}

func ObtenerUsuario(r *http.Request) modelos.Admin {
	var usuario modelos.Admin
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		log.Println("Error al decodificar JSON:", err)
	}
	return usuario
}
