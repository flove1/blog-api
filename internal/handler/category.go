package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCategories(ctx *gin.Context) {
	categories, err := h.Srvs.GetCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "ok",
		Body:    categories,
	})
}
