package model

import "time"

type Email struct {
	Recipients []string `json:"recipients"`
	Subject    string   `json:"subject"`
	Content    string   `json:"content"`
}

type EmailNotification struct {
	ID             int64      `json:"id"`
	UserID         int64      `json:"user_id"`
	RecipientEmail string     `json:"recipient_email"`
	EventType      string     `json:"event_type"`
	Subject        string     `json:"subject"`
	Body           string     `json:"body"`
	Status         string     `json:"status"`
	SentAt         *time.Time `json:"sent_at"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
