package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/Tilvaldiyev/blog-api/pkg/jwtutil"
	"github.com/gin-gonic/gin"
)

func authenticationMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "authorixation header is empty",
			})
			return
		}

		jwtToken := strings.Split(authHeader, " ")
		if len(jwtToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: "authorization header is malformed",
			})
			return
		}

		claims, err := jwtutil.ValidateToken(jwtToken[1], cfg.AUTH.SigningKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		err = jwtutil.ValidatePayload(claims)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		userID, err := strconv.Atoi(claims.Subject)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}
		c.Set("userID", int64(userID))
		c.Next()
	}
}
