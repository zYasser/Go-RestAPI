package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/zYasser/Go-RestAPI/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	APIKey string `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"URL"`
	UserID    uuid.UUID `json:"user_id"`
}


func databaseUserToUserModel(dbUser database.User) User{
	return User{
		ID:dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		APIKey: dbUser.ApiKey,
	}
}

func databaseFeedToFeedModel(feed database.Feed) Feed{
	return Feed{
		ID:feed.ID,
		CreatedAt: feed.CreatedAt,
		UpdatedAt: feed.UpdatedAt,
		Name: feed.Name,
		Url: feed.Url,
		UserID:feed.UserID,
	}
}
