package jwtutil

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired = errors.New("token is expired")
	ErrTokenInvalid = errors.New("provided token in invalid")
)

func ValidatePayload(claims *jwt.RegisteredClaims) error {
	if time.Now().After(claims.ExpiresAt.Local()) {
		return ErrTokenExpired
	}

	return nil
}

func GenerateToken(userID int64, expiresIn time.Duration, secretKey string) (string, error) {
	payload := jwt.RegisteredClaims{
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(secretKey))
}

func ValidateToken(token string, secretKey string) (*jwt.RegisteredClaims, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return nil, ErrTokenInvalid
	}

	payload, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, ErrTokenInvalid
	}

	return payload, nil
}
