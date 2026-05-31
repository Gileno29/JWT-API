package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	jwt.RegisteredClaims
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}
