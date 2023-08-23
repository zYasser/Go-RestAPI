package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString== ""{
		log.Fatal("PORT isn't available")
	}
	router:=chi.NewRouter()
	srv :=&http.Server{
		Handler:router,
		Addr: ":"+portString,
	}
	log.Printf("Server Running At Port %s" , portString)
	err :=srv.ListenAndServe()
	if err!=nil{
		log.Fatalf("Error %s" ,err)
	}
	
	
	fmt.Printf("Port %s is running" , portString)
}