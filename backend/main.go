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
	defer config.DB.Close()

	if err := config.CreateDatabase(); err != nil {
		log.Fatal("Erreur lors de l'initialisation de la base de données:", err)
	}

	r := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	// =============================================================
	//                          PUBLIC ROUTES
	// =============================================================

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// =============================================================
	//                          PROTECTED ROUTES
	// =============================================================

	r.HandleFunc("/me", middleware.AuthMiddleware(controllers.Me)).Methods("GET")
	r.HandleFunc("/upgrade", middleware.AuthMiddleware(controllers.Upgrade)).Methods("POST")
	r.HandleFunc("/pairs", middleware.AuthMiddleware(controllers.GetPairs)).Methods("GET")

	handler := corsMiddleware(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(srv.ListenAndServe())
}
