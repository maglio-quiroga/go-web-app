package modelos

import (
	"errors"
	"go-web-app/db"
	"go-web-app/herramientas"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Nombre string `gorm:"not null" json:"nombre"`
	Email  string `gorm:"not null;unique" json:"email"`
	Clave  string `gorm:"not null" json:"clave"`
}

var DefaultAdminServ ServAdmin

type ServAdmin struct{}

func (ServAdmin) AutenticarAdmin(w http.ResponseWriter, admin Admin) {
	var authAdmin Admin
	result := db.Database.Where("email = ?", admin.Email).First(&authAdmin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(authAdmin.Clave), []byte(admin.Clave))
	if err != nil {
		http.Error(w, "Clave incorrecta", http.StatusUnauthorized)
		return
	}
	token, err2 := herramientas.GenerarToken(32)
	if err2 != nil {
		http.Error(w, "Hubo un error al generar el token", http.StatusInternalServerError)
		return
	}
	expiracionToken := time.Now().Add(2 * time.Hour)
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expiracionToken,
		Path:    "/",
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Autenticación exitosa"))
}

func HashClaveAdm(clave string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(clave), bcrypt.DefaultCost)
	return string(hash), err
}
