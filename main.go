package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/nschimek/golang-http-server/internal/database"
)

type apiConfig struct {
	dbClient database.Client
}

const dbPath = "db.json"

func main() {
	m := http.NewServeMux()

	db := database.NewClient(dbPath)
	err := db.EnsureDB()
	if err != nil {
		log.Fatal(err)
	}

	api := apiConfig{dbClient: db}

	m.HandleFunc("/users", api.endpointUsersHandler)
	m.HandleFunc("/users/", api.endpointUsersHandler)
	m.HandleFunc("/posts", api.endpointPostsHandler)
	m.HandleFunc("/posts/", api.endpointPostsHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on:", addr)
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func getParamFromPath(path, endpoint string) (string, error) {
	if strings.HasPrefix(path, endpoint) {
		param := strings.TrimPrefix(path, endpoint)
		if param == "" {
			return "", errors.New("param value not found")
		} else {
			return param, nil
		}
	} else {
		return "", errors.New("expected URL prefix not found")
	}
}
