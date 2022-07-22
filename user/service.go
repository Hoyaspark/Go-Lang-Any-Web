package user

import (
	"anyweb/config"
	"anyweb/util"
	"context"
	"log"
)

func Login(ctx context.Context, param *LoginRequestBody) (*util.JwtToken, error) {

	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return nil, err
	}

	repo := NewUserRepository(ctx, nil, db)

	u, err := repo.findPasswordByEmail(param.Email)

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

func Join(ctx context.Context, param *JoinRequestBody) error {
	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return err
	}

	repo := NewUserRepository(ctx, nil, db)

	if _, err := repo.findByEmail(param.Email); err != nil {
		return err
	}

	if err := repo.InsertIntoUser(NewUser(param.Email, EncryptPassword(param.Password), param.Name, param.Gender)); err != nil {
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
