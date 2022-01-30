package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorBody struct {
	Error string `json:"error:"`
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	errorBody := errorBody{Error: err.Error()}

	respondWithJSON(w, code, errorBody)
}
