package modelos

import (
	"time"

	"gorm.io/gorm"
)

type Evento struct {
	gorm.Model
	Nombre    string    `gorm:"not null"`
	Ubicacion string    `gorm:"not null"`
	Imagen    string    `gorm:"not null"`
	Inicio    time.Time `gorm:"not null"`
	Final     time.Time `gorm:"not null"`
}
