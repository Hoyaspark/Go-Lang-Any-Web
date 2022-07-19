package user

import (
	"anyweb/config"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotSaveUserInContext = errors.New("not found user in context")
)

type User struct {
	Id       int64 // BIGINT
	Email    string
	Password string
	Name     string
	Gender   bool
}

func (u *User) UpdatePassword(ctx context.Context, password string) string {
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	encPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		cancel()
	}

	return string(encPwd)
}

func (u *User) EncryptPassword() string {

	enc, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	return string(enc)

}

func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, config.UserKey, u)
}

func FromContext(ctx context.Context) (*User, error) {
	if u, ok := ctx.Value(config.UserKey).(*User); ok {
		return u, nil
	} else {
		return nil, ErrNotSaveUserInContext
	}
}
