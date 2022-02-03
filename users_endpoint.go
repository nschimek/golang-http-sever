package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (api apiConfig) endpointUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.handlerGetUser(w, r)
	case http.MethodPost:
		api.handlerCreateUser(w, r)
	case http.MethodPut:
		api.handlerUpdateUser(w, r)
	case http.MethodDelete:
		api.handlerDeleteUser(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (api apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("creating new user...")
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

	respondWithJSON(w, http.StatusCreated, user)
}

func (api apiConfig) handlerDeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting user...")
	userEmail, err := getParamFromPath(r.URL.Path, "/users/")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = api.dbClient.DeleteUser(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (api apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	log.Println("getting user...")
	userEmail, err := getParamFromPath(r.URL.Path, "/users/")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := api.dbClient.GetUser(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (api apiConfig) handlerUpdateUser(w http.ResponseWriter, r *http.Request) {
	log.Println("updating user...")
	userEmail, err := getParamFromPath(r.URL.Path, "/users/")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}{}

	err = decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("could not decode input params: "+err.Error()))
		return
	}

	user, err := api.dbClient.UpdateUser(userEmail, params.Password, params.Name, params.Age)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}
