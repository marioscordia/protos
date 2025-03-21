package model

import "time"

type Note struct {
	ID            int64     `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	OwnerID       int64     `json:"owner_id"`
	IsPublic      bool      `json:"is_public"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Tags          []string  `json:"tags"`
	Collaborators []int64   `json:"collaborators"`
}

type Collaborator struct {
	UserID   int64 `json:"user_id"`
	Username int64 `json:"username"`
}
