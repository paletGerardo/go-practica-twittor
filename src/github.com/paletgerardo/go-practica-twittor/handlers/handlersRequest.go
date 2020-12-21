package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
)

func HandlerRequest () {
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal( http.ListenAndServe(":"+port, handler))
}
