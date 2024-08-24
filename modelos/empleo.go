package modelos

import (
	"time"

	"gorm.io/gorm"
)

type Empleo struct {
	gorm.Model
	Titulo      string    `gorm:"not null"`
	Descripcion string    `gorm:"not null"`
	Fecha       time.Time `gorm:"not null"`
	Empresa     string    `gorm:"not null"`
	Ciudad      string    `gorm:"not null"`
	Sueldo      string    `gorm:"not null"`
	Archivo     string    `gorm:"not null"`
}
