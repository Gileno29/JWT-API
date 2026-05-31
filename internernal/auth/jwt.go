package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secretKey []byte
}

func NewJWT(secretKey []byte) *JWT {
	return &JWT{secretKey: secretKey}
}

func (j *JWT) GenerateToken(userID uint64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		UserClaims{
			ID:    userID,
			Email: email,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		})

	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (j *JWT) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
