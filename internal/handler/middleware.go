package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("authorization header is missing")
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("bad signed method received")
		}
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func authenticate(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		token, err := parseToken(jwtToken, string(cfg.HTTP.SigningKey))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: "bad jwt token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: "unable to parse claims",
			})
			return
		}

		c.Set("userID", (int64)(claims["userID"].(float64)))
		c.Next()
	}
}
