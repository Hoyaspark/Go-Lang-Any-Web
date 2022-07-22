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

func Login(ctx context.Context, param *LoginRequestBody) (*util.JwtToken, error) {

	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return nil, err
	}

	repo := NewUserRepository(ctx, nil, db)

	u, err := repo.findPasswordByEmail(param)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := u.MatchPassword(param.Password); err != nil {
		log.Println(err)
		return nil, err
	}

	return util.GenerateJwtToken(param.Email)
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

	row, err := tx.Query("SELECT u.id FROM user AS u WHERE u.email = ?", u.email)
	row.Next()
	if err != nil {
		log.Println(err)
		return ErrDuplicateEmail
	}

	if _, err := tx.Exec("INSERT INTO user(email,password,name,gender) VALUES (?,?,?,?)", u.email, u.EncryptPassword(), u.name, u.gender); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

//func GetUserInfo(ctx context.Context, u *User) error {
//	db, err := config.DatabaseFromContext(ctx)
//
//	if err != nil {
//		return err
//	}
//
//	tx, err := db.Begin()
//
//	if err != nil {
//		return err
//	}
//
//	defer tx.Rollback()
//
//	tx.QueryRow("SELECT * FROM user as u WHERE u.email = ?", u.email).Scan()
//
//}
