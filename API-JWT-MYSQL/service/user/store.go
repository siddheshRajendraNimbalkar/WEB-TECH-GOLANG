package user

import (
	"database/sql"
	"fmt"

	types "github.com/siddheshRajendraNimbalkar/API-JWT-MYSQL/TYPES"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	row, err := s.db.Query("SELECT * FROM user WHERE email = ?", email)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user types.User
	if row.Next() {
		err := row.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("user not found")
	}

	return &user, nil
}

func (s *Store) CreateUser(user types.User) error {
	// Insert user into the database
	_, err := s.db.Exec(
		"INSERT INTO user (firstName, lastName, email, password, createdAt) VALUES (?, ?, ?, ?, ?)",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.CreatedAt,
	)
	return err
}
