package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT isn't available")
	}
	router := chi.NewRouter()
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	v1Router:=chi.NewRouter()
	v1Router.HandleFunc("/health" , handlerReadiness)
	router.Mount("/v1" , v1Router)
	log.Printf("Server Running At Port %s", portString)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Error %s", err)
	}
	
	
	fmt.Printf("Port %s is running", portString)
}
