package main

import (
	"fmt"
	"net/http"

	"github.com/prantoran/rssagg/internal/auth"
	"github.com/prantoran/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) authMiddleware(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respWithErr(w, http.StatusForbidden, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respWithErr(w, http.StatusBadRequest, fmt.Sprintf("Error fetching user: %v", err))
			return
		}

		next(w, r, user)
	}
}
