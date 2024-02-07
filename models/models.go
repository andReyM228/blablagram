package models

type Feed struct {
	User    FeedUser
	Posts   []Post
	Stories []FeedStory
}

// LoginUser is a model for user login
type LoginUser struct {
	Email    string
	Password string
}

// User is a model for user registration
type User struct {
	ID        string `bson:"_id"`
	Username  string `bson:"username"`
	FullName  string `bson:"full_name"`
	Email     string `bson:"email"`
	Password  string `bson:"password"`
	AvatarURL string `bson:"avatar_url"`
	CreatedAt string `bson:"created_at"`
	IsLogged  bool   `bson:"is_logged"`
}

type Post struct {
	ID          string        `bson:"_id"`
	AuthorID    string        `bson:"author_id"`
	ImageURL    string        `bson:"image_url"`
	VideoURL    string        `bson:"video_url"`
	Description string        `bson:"description"`
	CreatedAt   string        `bson:"created_at"`
	Comments    []FeedComment `bson:"comments"`
}

type FeedComment struct {
	ID        string `bson:"_id"`
	Author    User   `bson:"author"`
	Text      string `bson:"text"`
	CreatedAt string `bson:"created_at"`
}

type FeedUser struct {
	ID        string
	Username  string
	FullName  string
	AvatarURL string
}

type FeedStory struct {
	ID        string
	Author    User
	CreatedAt string
	ExpiresAt string
}
