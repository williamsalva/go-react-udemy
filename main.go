package main

import (
	"log"

	"github.com/ramirowilliams/go-react-udemy-back/bd"
	"github.com/ramirowilliams/go-react-udemy-back/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}