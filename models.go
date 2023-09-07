package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/zYasser/Go-RestAPI/internal/database"
)
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"URL"`
	UserID    uuid.UUID `json:"user_id"`
}
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	APIKey string `json:"api_key"`
}



type Follow_Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedId    uuid.UUID `json:"feed_id"`

}


type Follow_Feed_List struct {
	User 
	Feeds []Follow_Feed `json:"feeds"`
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
func databaseFeedsToFeedModel(feeds []database.Feed) []Feed {
	result := make([]Feed, len(feeds))

	for i, feed := range feeds {
		result[i] = Feed{
			ID:        feed.ID,
			CreatedAt: feed.CreatedAt,
			UpdatedAt: feed.UpdatedAt,
			Name:      feed.Name,
			Url:       feed.Url,
			UserID:    feed.UserID,
		}
	}

	return result
}
func databaseFollowFeedsToModel(followFeeds database.FeedFollow) Follow_Feed{
	return Follow_Feed{
		ID:followFeeds.ID,
		CreatedAt: followFeeds.CreatedAt,
		UpdatedAt: followFeeds.UpdatedAt,
		UserID:followFeeds.UserID,
		FeedId: followFeeds.FeedID,

	}
}
func databaseFollowFeedsToListModel(followFeeds []database.FeedFollow , user database.User) Follow_Feed_List{
	result := make([]Follow_Feed , len(followFeeds))
	for i , followFeed := range followFeeds {
		result[i]=Follow_Feed{
			ID:followFeed.ID,
			CreatedAt: followFeed.CreatedAt,
			UpdatedAt: followFeed.UpdatedAt,
			UserID:followFeed.UserID,
			FeedId: followFeed.FeedID,
	
		}
	}
	
	return Follow_Feed_List{
		User:databaseUserToUserModel(user),
		Feeds: result,

	} 
}
