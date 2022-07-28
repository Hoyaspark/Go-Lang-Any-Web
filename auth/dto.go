package auth

import (
	"errors"
	"strconv"
)

type InfoResponseBody struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Gender Gender `json:"gender"`
}

func NewInfoResponseBody(m *Member) *InfoResponseBody {
	return &InfoResponseBody{
		Email:  m.email,
		Name:   m.name,
		Gender: m.gender,
	}
}

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type JoinRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Gender   `json:"gender"`
}

type Gender struct {
	bool
}

func NewGender(b bool) Gender {
	return Gender{b}
}

func (g *Gender) Get() bool {
	return g.bool
}

func (g *Gender) MarshalJSON() ([]byte, error) {
	if g.bool == true {
		return []byte("\"" + "F" + "\""), nil
	} else if g.bool == false {
		return []byte("\"" + "M" + "\""), nil
	}

	return nil, errors.New("MarshalJSON Err")
}

func (g *Gender) UnmarshalJSON(d []byte) error {
	str := string(d)

	if str == "F" {
		g.bool = true
		return nil
	} else if str == "M" {
		g.bool = false
		return nil
	}

	return strconv.ErrSyntax
}
