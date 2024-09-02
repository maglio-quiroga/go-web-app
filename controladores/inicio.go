package controladores

import (
	"encoding/json"
	"fmt"
	"go-web-app/db"
	"go-web-app/modelos"
	"net/http"
	"sync"

	"gorm.io/gorm"
)

func InicioGet(w http.ResponseWriter, r *http.Request) {
	var alumnos []modelos.Alumno
	var eventos []modelos.Evento
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		db.Database.Order("egreso desc").Limit(3).Find(&alumnos)
		wg.Done()
	}()
	go func() {
		db.Database.Find(&eventos)
		wg.Done()
	}()
	wg.Wait()
	resp := map[string]interface{}{
		"Alumnos": alumnos,
		"Eventos": eventos,
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InicioPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	peticion, err := ObtenerPeticion(r)
	if err != nil { //falla la decodificacion a json
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err2 := GuardarPeticion(peticion)
	if err2 != nil { //falla la consulta sql
		http.Error(w, err2.Error(), http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Peticion Insertada"))
	}
}

func GuardarPeticion(p modelos.Peticion) error {
	var alumno modelos.Alumno
	err := db.Database.Where("email = ?", p.Email).First(&alumno).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if err == gorm.ErrRecordNotFound {
		return fmt.Errorf("el email %s no se encuentra en la tabla de alumnos", p.Email)
	}
	result := db.Database.Where("email = ?", p.Email).Assign(p).FirstOrCreate(&p)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func ObtenerPeticion(r *http.Request) (modelos.Peticion, error) {
	var peticiones modelos.Peticion
	err := json.NewDecoder(r.Body).Decode(&peticiones)
	if err != nil {
		return peticiones, err
	}
	return peticiones, nil
}
