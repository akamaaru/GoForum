package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type PostStore interface {
	GetPosts() ([]Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(Post) error
	// TODO DeletePostByID(id int) error
}

type User struct {
	ID        int
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterUserPayload struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=30"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type GetFeedPayload struct{}

type GetPostByIDPayload struct{}

type CreatePostPayload struct {
	Title string `json:"title" validate:"required,min=3,max=30"`
	Text  string `json:"text" validate:"required,min=3,max=100"`
}
