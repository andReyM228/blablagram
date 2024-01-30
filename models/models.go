package models

import "time"

type Feed struct {
	User    FeedUser
	Posts   []Post
	Stories []FeedStory
}

type User struct {
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
