package main

import (
	"log"
	"net/http"
	"sistemas/rutas"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/", rutas.Inicio)
	http.HandleFunc("/crear", rutas.Crear)
	http.HandleFunc("/insertar", rutas.Insertar)
	http.HandleFunc("/actualizar", rutas.Actualizar)
	http.HandleFunc("/editar", rutas.Editar)
	http.HandleFunc("/borrar", rutas.Borrar)

	log.Println("Servidor corriendo...")

	http.ListenAndServe(":8080", nil)

}
