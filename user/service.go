package user

import (
	"anyweb/config"
	"anyweb/util"
	"context"
	"errors"
	"log"
)

var (
	ErrNotFoundUser   = errors.New("not found user in database")
	ErrDuplicateEmail = errors.New("duplicate user")
)

func Login(ctx context.Context, u *User) (*util.JwtToken, error) {

	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return nil, err
	}

	tx, err := db.Begin()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer tx.Rollback()

	var pwd string

	if err := tx.QueryRow("SELECT U.password FROM user AS U WHERE U.email = ?", u.Email).Scan(&pwd); err != nil {
		log.Println(err)
		return nil, err
	}

	if err := u.MatchPassword(pwd); err != nil {
		log.Println(err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Println(err)
		return nil, err
	}

	return util.GenerateJwtToken(u.Email)
}

func Join(ctx context.Context, u *User) error {
	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return err
	}

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var count int
	if err := tx.QueryRow("SELECT count(*) FROM user AS u WHERE u.email = ?", u.Email).Scan(&count); err != nil || count != 0 {
		return ErrDuplicateEmail
	}

	if _, err := tx.Exec("INSERT INTO user(email,password,name,gender) VALUES (?,?,?,?)", u.Email, u.EncryptPassword(), u.Name, u.Gender); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
