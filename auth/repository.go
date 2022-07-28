package auth

import (
	"context"
	"database/sql"
	"errors"
	"sync"
)

var (
	ErrNotFoundUser   = errors.New("not found auth in database")
	ErrDuplicateEmail = errors.New("duplicate auth")
)

type Repository interface {
	findByEmail(email string) (*Member, error)
	findPasswordByEmail(email string) (*Member, error)
	insertIntoUser(user *Member) error
}

type MemberRepository struct {
	ctx context.Context
	mu  *sync.Mutex
	db  *sql.DB
}

func NewMemberRepository(ctx context.Context, mu *sync.Mutex, db *sql.DB) *MemberRepository {
	if mu == nil {
		return &MemberRepository{ctx: ctx, mu: &sync.Mutex{}, db: db}
	}
	return &MemberRepository{ctx: ctx, mu: mu, db: db}
}

func (mr *MemberRepository) findByEmail(email string) (*Member, error) {
	tx, err := mr.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var u Member

	if err := tx.QueryRow("SELECT u.email,u.password,u.name,u.gender FROM user AS u WHERE u.email = ?", email).Scan(&u.email, &u.password, &u.name, &u.gender); err == nil {
		return nil, ErrDuplicateEmail
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (mr *MemberRepository) findPasswordByEmail(email string) (*Member, error) {
	tx, err := mr.db.Begin()

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	var u Member

	if err := tx.QueryRow("SELECT u.password FROM user as u WHERE u.email = ?", email).Scan(&u.password); err != nil {
		return nil, ErrNotFoundUser
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (mr *MemberRepository) insertIntoUser(u *Member) error {
	tx, err := mr.db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	if err := tx.QueryRow("INSERT INTO auth(email,password,name,gender) VALUES (?,?,?,?)", u.email, u.password, u.name, u.gender).Err(); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
