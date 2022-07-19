package util

import (
	"anyweb/config"
	"anyweb/user"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtToken struct {
	Type         string    `json:"type"`
	AcceptToken  string    `json:"acceptToken"`
	RefreshToken string    `json:"refreshToken"`
	ExpiredAt    time.Time `json:"expiredAt"`
}

func GenerateJwtToken(email string) (*JwtToken, error) {
	exp := time.Now().Add(time.Minute * 30)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   exp,
	})

	tokenString, err := token.SignedString([]byte(config.AuthProperties.JwtSecret))

	if err != nil {
		return nil, err
	}

	return &JwtToken{Type: "Bearer", AcceptToken: tokenString, ExpiredAt: exp}, nil

}

func ParseJwtToken(token string) (*user.User, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.AuthProperties.JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	m := t.Claims.(jwt.MapClaims)

	return &user.User{Email: m["userEmail"].(string)}, nil
}
