package post

import (
	"database/sql"
	"fmt"

	"github.com/akamaaru/go-forum/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetPosts() ([]types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	posts := []types.Post {}
	for rows.Next() {
		post, err := scanRowToPost(rows)
		if err != nil {
			return nil, err
		}

		posts = append(posts, *post)
	}

	return posts, nil
}

func (s *Store) GetPostByID(id int) (*types.Post, error) {
	rows, err := s.db.Query("SELECT * FROM posts WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	post := new(types.Post)
	for rows.Next() {
		post, err = scanRowToPost(rows)
		if err != nil {
			return nil, err
		}
	}

	if post.ID == 0 {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (s *Store) CreatePost(post types.Post) error {
	_, err := s.db.Exec(
		"INSERT INTO posts (user_id, title, text) VALUES (?,?,?)", 
		post.UserID, post.Title, post.Text,
	)

	return err
}

func scanRowToPost(rows *sql.Rows) (*types.Post, error) {
	post := new(types.Post)

	err := rows.Scan(
		&post.ID,
		&post.UserID,
		&post.Title,
		&post.Text,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return post, nil
}