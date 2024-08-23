package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn string = "host=localhost user=postgres password=maglio dbname=WebAlumnos port=5432"
var Database *gorm.DB

func Iniciar_conexion() {
	var err error
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("CONEXION EXITOSA")
	}
}
