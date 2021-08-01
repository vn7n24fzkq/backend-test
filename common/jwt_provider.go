package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type jwtUserClaims struct {
	User JWTUser `json:"user"`
	jwt.StandardClaims
}

// Make sure this secret is private in your project
var hmacSecret string = "casper"

func GetJWT(jwtUser JWTUser) (string, error) {
	// Create the Claims
	claims := jwtUserClaims{
		jwtUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // expired after 24 hours
			Issuer:    "vn7n24fzkq",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(hmacSecret))
	return tokenString, err
}

func DecodeJWT(tokenString string) (JWTUser, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwtUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(hmacSecret), nil
	})

	if err != nil {
		return JWTUser{}, err
	}

	claims, ok := token.Claims.(*jwtUserClaims)
	if ok && token.Valid {
		return claims.User, nil
	} else {
		return claims.User, errors.New("JWT is not valid")
	}
}
