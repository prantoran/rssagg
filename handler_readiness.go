package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respWithJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}
