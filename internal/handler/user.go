package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
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

	user, err := h.Srvs.GetUserByUsername(ctx, req.Username)
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

	token, err := h.Srvs.CreateToken(ctx, user, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			http.StatusBadRequest,
			err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "token succesfully created",
		Body:    LoginResponseBody{token},
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
