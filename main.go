package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nschimek/golang-http-server/internal/database"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", testHandler)
	m.HandleFunc("/err", testErrHandler)

	const addr = "localhost:8080"
	srv := http.Server{
		Handler:      m,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on:", addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, database.User{
		Email: "test@example.com",
	})
}

func testErrHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, errors.New("test error"))
}
