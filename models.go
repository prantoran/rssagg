package main

import (
	"github.com/google/uuid"
	"github.com/prantoran/rssagg/internal/database"
)

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	APIKey    string `json:"api_key,omitempty"`
}

type Feed struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID.String(),
		Name:      dbUser.Name,
		CreatedAt: dbUser.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: dbUser.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		APIKey:    dbUser.ApiKey,
	}
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID.String(),
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		CreatedAt: dbFeed.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: dbFeed.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UserID:    dbFeed.UserID,
	}
}
