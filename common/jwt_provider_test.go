package common

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAndDecodeCorrectJWT(t *testing.T) {
	jwtUser := JWTUser{
		1,
		"testingUser",
	}

	token, getJWTError := GetJWT(jwtUser)
	if getJWTError != nil {
		t.Fatal("Should not get error when GetJWT")
	}

	decodedJWTUser, decodeJWTError := DecodeJWT(token)
	if decodeJWTError != nil {
		t.Fatal("Should not get error when DecodeJWT")
	}

	if !cmp.Equal(jwtUser, decodedJWTUser) {
		t.Fatal("JWTUser should not be changed")
	}
}

func TestGetAndDecodeIncorrectJWT(t *testing.T) {
	jwtUser := JWTUser{
		1,
		"testingUser",
	}

	token, getJWTError := GetJWT(jwtUser)
	if getJWTError != nil {
		t.Fatal("Should not get error when GetJWT")
	}

	_, decodeJWTError := DecodeJWT(token + "wrongstring")
	if decodeJWTError == nil {
		t.Fatal("Should get an error when decode wrong JWT")
	}
}
