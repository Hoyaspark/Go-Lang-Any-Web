package user

import (
	"context"
	"database/sql"
	"sync"
)

type repository interface {
	findByEmail()
	findPasswordByEmail(param *LoginRequestBody) (*User, error)
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

func (ur *userRepository) findByEmail() {

}

func (ur *userRepository) findPasswordByEmail(param *LoginRequestBody) (*User, error) {
	tx, err := ur.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var u *User

	if err := tx.QueryRow("SELECT u.password FROM user as u WHERE u.email = ?", param.Email).Scan(u.password); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return u, nil
}
