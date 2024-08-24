package modelos

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Nombre string `gorm:"not null"`
	Email  string `gorm:"not null"`
	Clave  string `gorm:"not null"`
}
