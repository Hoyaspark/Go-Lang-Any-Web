package auth

import (
	"anyweb/config"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotSaveUserInContext = errors.New("not found auth in context")
)

type Member struct {
	id       int64 // BIGINT
	email    string
	password string
	name     string
	gender   bool
}

func NewMember(email, password, name string, gender bool) *Member {
	return &Member{email: email, password: password, name: name, gender: gender}
}

func (u *Member) Email() string {
	return u.email
}

func (u *Member) Password() string {
	return u.password
}

func (u *Member) UpdatePassword(ctx context.Context, password string) string {
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	encPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		cancel()
	}

	return string(encPwd)
}

func (u *Member) MatchPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
}

func EncryptPassword(password string) string {

	enc, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(enc)

}

func ContextWithMember(ctx context.Context, u *Member) context.Context {
	return context.WithValue(ctx, config.MemberKey, u)
}

func MemberFromContext(ctx context.Context) (*Member, error) {
	if u, ok := ctx.Value(config.MemberKey).(*Member); ok {
		return u, nil
	} else {
		return nil, ErrNotSaveUserInContext
	}
}
