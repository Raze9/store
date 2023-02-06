package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id            uint   `json:"id"`
	UserName      string `json:"user_Name"`
	Authorization int    `json:"authorization"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, Authorization int) (string, error) {
	notTime := time.Now()
	expireTime := notTime.Add(24 * time.Hour)
	claims := Claims{
		Id:            id,
		UserName:      username,
		Authorization: Authorization,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "mail",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
