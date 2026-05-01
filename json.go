package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5XX err: %v", msg)
	}

	type errResp struct {
		Error string `json:"error"`
	}

	respWithJSON(w, code, errResp{
		Error: msg,
	})
}

func respWithJSON(w http.ResponseWriter, code int, data any) {
	dat, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
