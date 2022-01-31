package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (api apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		api.handlerCreateUser(w, r)
	case http.MethodDelete:
		api.handlerDeleteUser(w, r)
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

func (api apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting user...")
	userEmail := strings.TrimPrefix(r.URL.Path, "/users/")

	if userEmail == "" {
		respondWithError(w, http.StatusBadRequest, errors.New("no email address specified in path"))
		return
	}

	err := api.dbClient.DeleteUser(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
