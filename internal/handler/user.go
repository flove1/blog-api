package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// var passwordRegEx = regexp.MustCompile(`^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{4,8}$`)

func (h *Handler) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			http.StatusBadRequest,
			err.Error(),
		})
		return
	}

	// if !passwordRegEx.Match([]byte(req.Password)) {
	// 	ctx.JSON(http.StatusBadRequest, &ErrorResponse{
	// 		http.StatusBadRequest,
	// 		"password is not strong enough",
	// 	})
	// }

	err = h.Srvs.CreateUser(ctx, &entity.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusCreated,
		Message: "user succesfully created",
	})
}

func (h *Handler) login(ctx *gin.Context) {
	var req LoginRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			http.StatusBadRequest,
			err.Error(),
		})
		return
	}

	user, err := h.Srvs.Login(ctx, req.Username, req.Password)
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			ctx.JSON(http.StatusBadRequest, &ErrorResponse{
				http.StatusBadRequest,
				err.Error(),
			})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
				http.StatusInternalServerError,
				err.Error(),
			})
			return
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    jwt.NewNumericDate(time.Now().Add(120 * time.Minute)),
	})

	tokenString, err := token.SignedString([]byte(h.Config.HTTP.SigningKey))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "token succesfully created",
		Body:    LoginResponseBody{tokenString},
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {
	var req UpdateUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			http.StatusBadRequest,
			err.Error(),
		})
		return
	}

	// if !passwordRegEx.Match([]byte(req.Password)) {
	// 	ctx.JSON(http.StatusBadRequest, &ErrorResponse{
	// 		http.StatusBadRequest,
	// 		"password is not strong enough",
	// 	})
	// }

	userID := ctx.MustGet("userID").(int64)

	err = h.Srvs.UpdateUser(ctx, &entity.User{
		ID:        userID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "user succesfully updated",
	})

}

func (h *Handler) deleteUser(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(int64)

	err := h.Srvs.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "user succesfully deleted",
	})

}
