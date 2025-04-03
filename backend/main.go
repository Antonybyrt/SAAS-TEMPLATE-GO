package main

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	if err := config.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Configuration CORS
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	// Routes publiques
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Routes protégées
	r.HandleFunc("/me", middleware.AuthMiddleware(controllers.Me)).Methods("GET")

	// Appliquer le middleware CORS
	handler := corsMiddleware(r)

	// Configuration du serveur
	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(srv.ListenAndServe())
}
