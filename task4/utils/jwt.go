package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type JwtCustClaims struct {
	ID   uint
	Name string
	jwt.RegisteredClaims
}

var stSigningKey = []byte(viper.GetString("jwt.signingKey"))

// 生成token
func GenerateToken(id uint, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(stSigningKey)
}

// 解析token
func ParseToken(tokenStr string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &iJwtCustClaims, func(token *jwt.Token) (interface{}, error) {
		return stSigningKey, nil
	})

	if err == nil && !token.Valid {
		err = errors.New("Invalid Token")
	}

	return iJwtCustClaims, err
}

func IsTokenValid(tokenStr string) bool {
	_, err := ParseToken(tokenStr)
	if err != nil {
		return false
	}

	return true
}
