package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

type IAuthRepo interface {
	GenerateJWT(userId int, email, password string) (string, error)
	ClaimToken(token string) (map[string]interface{}, error)
}

type authRepo struct {
	jwt          *jwt.Token
	signingKey   []byte
	expiredToken int
}

func (a authRepo) GenerateJWT(userId int, email, password string) (string, error) {
	claims := a.jwt.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user_id"] = strconv.Itoa(userId)
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(a.expiredToken)).Unix()
	return a.jwt.SignedString(a.signingKey)
}

func (a authRepo) ClaimToken(token string) (map[string]interface{}, error) {
	tempToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return a.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	claim, ok := tempToken.Claims.(jwt.MapClaims)
	if ok && tempToken.Valid {
		return claim, nil
	}
	return nil, fmt.Errorf("not Authorize")
}

func NewAuthRepo(jwt *jwt.Token, signingKey []byte, expiredToken int) IAuthRepo {
	return &authRepo{jwt: jwt, signingKey: signingKey, expiredToken: expiredToken}
}
