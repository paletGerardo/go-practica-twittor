package main

import (
	"github.com/paletgerardo/twittor/handlers"
	"github.com/paletgerardo/twittor/bd"
	"log"
)

func main() {

	if bd.CheckConection() == 0 {
		log.Fatal("Error al conectar con bd")
		return
	}

	handlers.HandlerRequest()

}
