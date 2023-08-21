package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load(".env")
	portString:=os.Getenv("PORT")
	if portString== ""{
		log.Fatal("PORT isn't available")
	}
	
	
	fmt.Printf("Port %s is running" , portString)
}