package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HOLA MUNDO")
}

func main() {
	enrutador := mux.NewRouter()
	enrutador.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	enrutador.HandleFunc("/", Index)
	http.ListenAndServe(":3000", enrutador)
}
