package user

import (
	"context"
	"database/sql"
	"errors"
	"sync"
)

var (
	ErrNotFoundUser   = errors.New("not found user in database")
	ErrDuplicateEmail = errors.New("duplicate user")
)

type repository interface {
	findByEmail(email string) (*User, error)
	findPasswordByEmail(email string) (*User, error)
	InsertIntoUser(user *User) error
}

type userRepository struct {
	ctx context.Context
	mu  *sync.Mutex
	db  *sql.DB
}

func NewUserRepository(ctx context.Context, mu *sync.Mutex, db *sql.DB) *userRepository {
	if mu == nil {
		return &userRepository{ctx: ctx, mu: &sync.Mutex{}, db: db}
	}
	return &userRepository{ctx: ctx, mu: mu, db: db}
}

func (ur *userRepository) findByEmail(email string) (*User, error) {
	tx, err := ur.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var u User

	if err := tx.QueryRow("SELECT u.email,u.password,u.name,u.gender FROM user AS u WHERE u.email = ?", email).Scan(&u.email, &u.password, &u.name, &u.gender); err == nil {
		return nil, ErrDuplicateEmail
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (ur *userRepository) findPasswordByEmail(email string) (*User, error) {
	tx, err := ur.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var u User

	if err := tx.QueryRow("SELECT u.password FROM user as u WHERE u.email = ?", email).Scan(&u.password); err != nil {
		return nil, ErrNotFoundUser
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (ur *userRepository) InsertIntoUser(u *User) error {
	tx, err := ur.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := tx.QueryRow("INSERT INTO user(email,password,name,gender) VALUES (?,?,?,?)", u.email, u.password, u.name, u.gender).Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
