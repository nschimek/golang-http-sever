package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/nschimek/golang-http-server/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

func (api apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		api.handlerCreateUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (api apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating new user...")
	decoder := json.NewDecoder(r.Body)
	params := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("could not decode input params: "+err.Error()))
		return
	}

	user, err := api.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, 201, user)
}
