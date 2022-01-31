package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nschimek/golang-http-server/internal/database"
)

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
