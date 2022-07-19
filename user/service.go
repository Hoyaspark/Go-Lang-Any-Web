package user

import (
	"anyweb/config"
	"anyweb/util"
	"context"
	"errors"
	"log"
)

var (
	ErrNotFoundUser = errors.New("not found user in database")
)

func Login(ctx context.Context, u *User) (*util.JwtToken, error) {
	context, cancel := context.WithCancel(ctx)

	defer cancel()

	db, err := config.DatabaseFromContext(context)

	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer tx.Rollback()

	if err := tx.QueryRow("SELECT * FROM user AS U WHERE U.email = ? AND U.password = ?", u.Email, u.EncryptPassword()).Scan(); err != nil {
		log.Println(err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil, err
	}

	return util.GenerateJwtToken(u.Email)
}
