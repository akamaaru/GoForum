package comment

import (
	"database/sql"

	"github.com/akamaaru/go-forum/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCommentsByPostID(id int) ([]types.Comment, error) {
	rows, err := s.db.Query("SELECT * FROM comments WHERE post_id = ?", id)
	if err != nil {
		return nil, err
	}

	comments := []types.Comment {}
	for rows.Next() {
		comment, err := scanRowToComment(rows)
		if err != nil {
			return nil, err
		}

		comments = append(comments, *comment)
	}

	return comments, nil
}

func (s *Store) CreateComment(comment types.Comment) error {
	_, err := s.db.Exec(
		"INSERT INTO comments (post_id, user_id, text) VALUES (?,?,?)", 
		comment.PostID, comment.UserID, comment.Text,
	)

	return err
}

func scanRowToComment(rows *sql.Rows) (*types.Comment, error) {
	comment := new(types.Comment)

	err := rows.Scan(
		&comment.ID,
		&comment.PostID,
		&comment.UserID,
		&comment.Text,
		&comment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return comment, nil
}