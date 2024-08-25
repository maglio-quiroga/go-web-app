package modelos

import (
	"errors"
	"go-web-app/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Nombre string `gorm:"not null"`
	Email  string `gorm:"not null;unique"`
	Clave  string `gorm:"not null"`
}

type Adm struct {
	Nombre string
	Email  string
	Clave  string
}

type AuthAdmin struct {
	Nombre    string
	Email     string
	ClaveHash string
}

var DefaultAdminServ ServAdmin

type ServAdmin struct{}

func (ServAdmin) AutenticarAdmin(admin Adm) error {
	var authAdmin AuthAdmin
	result := db.Database.Where("email = ?", admin.Email).First(&authAdmin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("Usuario no encontrado")
	}
	err := bcrypt.CompareHashAndPassword([]byte(authAdmin.ClaveHash), []byte(admin.Clave))
	if err != nil {
		return errors.New("Clave incorrecta")
	}
	return nil
}

func HashClaveAdm(clave string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(clave), bcrypt.DefaultCost)
	return string(hash), err
}
