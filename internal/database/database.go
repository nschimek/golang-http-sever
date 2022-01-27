package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Client struct {
	path string
}

func NewClient(path string) Client {
	return Client{path: path}
}

type DatabaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func (c Client) createDB() error {
	data, err := json.Marshal(DatabaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})

	if err != nil {
		return err
	}

	fmt.Println(data)
	err = os.WriteFile(c.path, data, 0640)

	if err != nil {
		return err
	}

	return nil
}

func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return c.createDB()
		}
		return err
	}

	return nil
}

func (c Client) updateDB(db DatabaseSchema) error {
	data, err := json.Marshal(db)

	if err != nil {
		return err
	}

	err = os.WriteFile(c.path, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (c Client) readDB() (DatabaseSchema, error) {
	data, err := os.ReadFile(c.path)

	if err != nil {
		return DatabaseSchema{}, err
	}

	db := DatabaseSchema{}
	err = json.Unmarshal(data, &db)

	return db, err
}
