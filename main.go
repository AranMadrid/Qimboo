package main

import (
	"log"

	"github.com/AranMadrid/Qimboo/bd"
	"github.com/AranMadrid/Qimboo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
