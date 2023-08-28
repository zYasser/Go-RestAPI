package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zYasser/Go-RestAPI/internal/auth"
	"github.com/zYasser/Go-RestAPI/internal/database"
)

func  (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Name string `json:"name"`
	}
	decoder :=json.NewDecoder(r.Body)
	params:=parameters{}
	err :=decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %s", err))
		return 
	}
	user , err := apiConfig.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err !=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user %s", err))

	}
	respondWithJSON(w, http.StatusCreated, databaseUserToUserModel(user)	)
}



func  (apiConfig *apiConfig)handlerGetUserByApiKey(w http.ResponseWriter, r *http.Request) {
	apiKey, err :=auth.GetAPIKeyFromHeader(r.Header)
	if err !=nil{
		respondWithError(w,403,fmt.Sprintf("Auth Error %s", err))
		return

	}
	user , err :=apiConfig.DB.GetUserByAPIKey(r.Context(),apiKey)
	if err !=nil{
	respondWithError(w,404,fmt.Sprintf("User doesn't exist %s", err))
	return

}
respondWithJSON(w,200,databaseUserToUserModel(user))

}
