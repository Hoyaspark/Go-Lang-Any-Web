package auth

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

	repo := NewMemberRepository(ctx, nil, db)

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

	repo := NewMemberRepository(ctx, nil, db)

	if _, err := repo.findByEmail(param.Email); err != nil {
		return err
	}

	if err := repo.insertIntoUser(NewMember(param.Email, EncryptPassword(param.Password), param.Name, param.Gender)); err != nil {
		return err
	}

	return nil
}

func GetUserInfo(ctx context.Context, m *Member) (*InfoResponseBody, error) {
	db, err := config.DatabaseFromContext(ctx)

	if err != nil {
		return nil, err
	}

	repo := NewMemberRepository(ctx, nil, db)

	u, err := repo.findByEmail(m.Email())

	if err != nil && err != ErrDuplicateEmail {
		return nil, err
	}

	return NewInfoResponseBody(u), nil

}
