package modelos

import (
	"gorm.io/gorm"
)

type Peticion struct {
	gorm.Model
	Nombre      string `gorm:"not null"`
	Email       string `gorm:"not null"`
	Area        string `gorm:"not null"`
	Trabajo     string `gorm:"not null"`
	Contacto    string `gorm:"not null"`
	Descripcion string `gorm:"not null"`
	Imagen      string `gorm:"not null"`
}
