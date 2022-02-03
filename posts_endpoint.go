package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func (api apiConfig) endpointPostsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.handlerGetPosts(w, r)
	case http.MethodPost:
		api.handlerCreatePost(w, r)
	case http.MethodDelete:
		api.handlerDeletePost(w, r)
	default:
		respondWithError(w, 404, errors.New("method not supported"))
	}
}

func (api apiConfig) handlerCreatePost(w http.ResponseWriter, r *http.Request) {
	log.Println("creating new post...")
	decoder := json.NewDecoder(r.Body)
	params := struct {
		UserEmail string `json:"userEmail"`
		Text      string `json:"text"`
	}{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, errors.New("could not decode input params:"+err.Error()))
		return
	}

	post, err := api.dbClient.CreatePost(params.UserEmail, params.Text)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, post)
}

func (api apiConfig) handlerDeletePost(w http.ResponseWriter, r *http.Request) {
	log.Println("deleting post...")
	postId, err := getParamFromPath(r.URL.Path, "/posts/")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	err = api.dbClient.DeletePost(postId)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

func (api apiConfig) handlerGetPosts(w http.ResponseWriter, r *http.Request) {
	log.Println("getting posts...")
	userEmail, err := getParamFromPath(r.URL.Path, "/posts/")

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	posts, err := api.dbClient.GetPosts(userEmail)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}
