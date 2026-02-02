package models

import (
	"time"
)

type Message struct {
	ID        string    `json:"id"`
	ChannelID string    `json:"channel_id"`
	UserID    string    `json:"user_id"`
	UserName  string    `json:"user_name"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`      // "chat", "join", "leave", "error"
	ParentID  string    `json:"parent_id"` // For replies
	CreatedAt time.Time `json:"created_at"`
}
