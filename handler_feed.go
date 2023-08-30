package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zYasser/Go-RestAPI/internal/database"
)

func  (apiConfig *apiConfig)handlerCreateFeed(w http.ResponseWriter, r *http.Request , user database.User ) {
	type parameters struct{
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder :=json.NewDecoder(r.Body)
	params:=parameters{}
	err :=decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %s", err))
		return 
	}
	feed , err := apiConfig.DB.CreateFeed(r.Context(),database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		UserID: user.ID,
		Url: params.URL,
	})
	if err !=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user %s", err))

	}
	respondWithJSON(w, http.StatusCreated, databaseFeedToFeedModel(feed))
}


