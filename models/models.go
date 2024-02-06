package models

import "time"

type Feed struct {
	User    FeedUser
	Posts   []Post
	Stories []FeedStory
}

// User is a model for user registration
type User struct {
	ID        string `json:"_id"`
	Username  string `json:"username"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarURL string `json:"avatar_url"`
	CreatedAt string `json:"created_at"`
	IsLogged  bool   `json:"is_logged"`
}

// LoginUser is a model for user login
type LoginUser struct {
	Email    string
	Password string
}

type FeedUser struct {
	ID        int64
	Username  string
	FullName  string
	AvatarURL string
}

type Post struct {
	ID          int64
	Author      User
	CreatedAt   time.Time
	ImageURL    string
	VideoURL    string
	Description string
}

type FeedComment struct {
	ID        int64
	Author    User
	Text      string
	CreatedAt time.Time
}

type FeedStory struct {
	ID        int64
	Author    User
	CreatedAt time.Time
	ExpiresAt time.Time
}
