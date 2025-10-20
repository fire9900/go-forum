package entity

import "time"

type Post struct {
	ID       int       `json:"id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at"`
	ThreadID int       `json:"thread_id"`
	UserID   int       `json:"user_id"`
	ParentID int       `json:"parent_id"`
}
