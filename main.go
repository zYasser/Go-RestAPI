package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/zYasser/Go-RestAPI/internal/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)
type apiConfig struct{
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT isn't available")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL isn't available")
	}
	conn , err :=sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Can't connect to database: " , err)
	}


	apiCfg := apiConfig{
		DB: database.New(conn),
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
	v1Router.Get("/health" , handlerReadiness)
	v1Router.Get("/error" , handlerErr)
	v1Router.Post("/users" , apiCfg.handlerCreateUser)
	v1Router.Get("/users" , apiCfg.middlewareAuth(apiCfg.handlerGetUserByApiKey))
	v1Router.Post("/feed" , apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feed" , apiCfg.handlerGetAllFeeds)
	v1Router.Post("/feed_follows" , apiCfg.middlewareAuth(apiCfg.handlerFollowFeed))
	v1Router.Get("/feed_follows" , apiCfg.middlewareAuth(apiCfg.handlerGetFollowedFeedByUser))


	router.Mount("/v1" , v1Router)
	log.Printf("Server Running At Port %s", portString)
	log.Fatal(srv.ListenAndServe())	
		
}
