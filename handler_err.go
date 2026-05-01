package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respWithErr(w, http.StatusInternalServerError, "This is an error response")
}
