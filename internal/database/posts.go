package database

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}
	_, err = c.GetUser(userEmail)
	if err != nil {
		return Post{}, err
	}

	post := Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UserEmail: userEmail,
		Text:      text,
	}

	db.Posts[post.ID] = post

	err = c.updateDB(db)

	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return nil, err
	}

	posts := make([]Post, 0)

	for _, v := range db.Posts {
		if v.UserEmail == userEmail {
			posts = append(posts, v)
		}
	}

	return posts, nil
}

func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}

	delete(db.Posts, id)

	err = c.updateDB(db)

	if err != nil {
		return err
	}

	return nil
}
