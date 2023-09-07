package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/zYasser/Go-RestAPI/internal/database"
)

func  (apiConfig *apiConfig)handlerFollowFeed(w http.ResponseWriter, r *http.Request , user database.User ) {
	type parameters struct{
		Feed_id uuid.UUID `json:"feed_id"`
	}
	decoder :=json.NewDecoder(r.Body)
	params:=parameters{}
	err :=decoder.Decode(&params)
	if err != nil{
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON %s", err))
		return 
	}
	feed , err := apiConfig.DB.CreateFeedFollow(r.Context(),database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.Feed_id,
		
	})
	if err !=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user %s", err))

	}
	respondWithJSON(w, http.StatusCreated, databaseFollowFeedsToModel(feed))
}





func  (apiConfig *apiConfig)handlerGetFollowedFeedByUser(w http.ResponseWriter, r *http.Request , user database.User ) {

	follow_feeds , err := apiConfig.DB.GetFeedFollowsByUser(r.Context(),user.ID)
	if err !=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't create user %s", err))

	}
	respondWithJSON(w, http.StatusOK, databaseFollowFeedsToListModel(follow_feeds ,user))
}

