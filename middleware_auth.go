package main

import (
	"fmt"
	"net/http"

	"github.com/zYasser/Go-RestAPI/internal/auth"
	"github.com/zYasser/Go-RestAPI/internal/database"
)

type authHandler func(http.ResponseWriter , *http.Request ,database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
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
		handler(w,r,user)
	}
}