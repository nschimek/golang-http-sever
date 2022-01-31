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
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)

	if err != nil {
		log.Println("error marshalling:", err)
		w.WriteHeader(500)
		response, _ := json.Marshal(errorBody{Error: "error marshalling"})
		w.Write(response)
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, err error) {
	log.Println(err)

	respondWithJSON(w, code, errorBody{Error: err.Error()})
}
