package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/prantoran/rssagg/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	params := parameter{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		respWithErr(w, http.StatusBadRequest, fmt.Sprintf("Error decoding request body: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respWithErr(w, http.StatusInternalServerError, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respWithJSON(w, http.StatusCreated, databaseFeedToFeed(feed))
}
