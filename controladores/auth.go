package controladores

import (
	"fmt"
	"go-web-app/modelos"
	"net/http"
)

func LogAdmin(w http.ResponseWriter, r *http.Request) {
	usuario := ObtenerUsuario(r)
	err := modelos.DefaultAdminServ.AutenticarAdmin(usuario)
	if err != nil {
		http.Error(w, "Credenciales incorrectas", http.StatusUnauthorized)
		return
	}
	fmt.Println("Iniciada la sesion")
}

func ObtenerUsuario(r *http.Request) modelos.Adm {
	nombre := r.FormValue("nombre")
	email := r.FormValue("email")
	clave := r.FormValue("clave")
	return modelos.Adm{
		Nombre: nombre,
		Email:  email,
		Clave:  clave,
	}
}
