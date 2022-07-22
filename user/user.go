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
	id       int64 // BIGINT
	email    string
	password string
	name     string
	gender   bool
}

func NewUser(email, password, name string, gender bool) *User {
	return &User{email: email, password: password, name: name, gender: gender}
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Password() string {
	return u.password
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

func (u *User) MatchPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.password), []byte(password))
}
func EncryptPassword(password string) string {

	enc, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(enc)

}

func ContextWithUser(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, config.UserKey, u)
}

func UserFromContext(ctx context.Context) (*User, error) {
	if u, ok := ctx.Value(config.UserKey).(*User); ok {
		return u, nil
	} else {
		return nil, ErrNotSaveUserInContext
	}
}
