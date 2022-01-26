package database

import "time"

type Client struct {
	path string
}

func NewClient(path string) Client {
	return Client{path: path}
}

type DatabaseSchema struct {
	users map[string]User `json:"users"`
	posts map[string]Post `json:"posts"`
}

type User struct {
	createdAt time.Time `json:"createdAt"`
	email     string    `json:"email"`
	password  string    `json:"password"`
	name      string    `json:"name"`
	age       int       `json:"age"`
}

type Post struct {
	id        string    `json:"id"`
	createdAt time.Time `json:"createdAt"`
	userEmail string    `json:"userEmail"`
	text      string    `json:"text"`
}
