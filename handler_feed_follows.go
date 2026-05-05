package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prantoran/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := parameter{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respWithErr(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respWithErr(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed follow: %v", err))
		return
	}

	respWithJSON(w, http.StatusCreated, databaseFeedFollowToFeedFollow(feedFollow))
}
