package modelos

import (
	"time"

	"gorm.io/gorm"
)

type Alumno struct {
	gorm.Model
	Nombre      string    `gorm:"not null"`
	Email       string    `gorm:"not null"`
	Area        string    `gorm:"not null"`
	Egreso      time.Time `gorm:"not null"`
	Trabajo     string    `gorm:"not null"`
	Contacto    string    `gorm:"not null"`
	Descripcion string    `gorm:"not null"`
	Imagen      string
}
