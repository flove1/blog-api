package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createArticle(ctx *gin.Context) {
	var req CreateArticleRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(int64)

	categories := make([]entity.Category, 0)
	for _, category := range req.Categories {
		categories = append(categories, entity.Category{
			Name: category,
		})
	}

	err = h.Srvs.CreateArticle(ctx, &entity.Article{
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Categories:  categories,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &Response{
		Code:    http.StatusCreated,
		Message: "article succesfully created",
	})
}

func (h *Handler) getArticleById(ctx *gin.Context) {
	var id ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	article, err := h.Srvs.GetArticleByID(ctx, id.Value)
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
		Body:    article,
	})
}

func (h *Handler) getAllArticles(ctx *gin.Context) {
	articles, err := h.Srvs.GetAllArticles(ctx)
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
		Body:    articles,
	})
}

func (h *Handler) getArticlesByUsername(ctx *gin.Context) {
	var username Username

	err := ctx.ShouldBindUri(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	articles, err := h.Srvs.GetArticlesByUsername(ctx, strings.ToLower(username.Value))
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
		Body:    articles,
	})
}

func (h *Handler) updateArticle(ctx *gin.Context) {
	var req UpdateArticleRequest

	err := ctx.ShouldBindUri(&req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("bind json err: %s \n", err.Error())
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(int64)

	categories := make([]entity.Category, 0)
	for _, category := range req.Categories {
		categories = append(categories, entity.Category{
			Name: category,
		})
	}

	err = h.Srvs.UpdateArticle(ctx, &entity.Article{
		ID:          req.ID.Value,
		Title:       req.Title,
		Description: req.Description,
		UserID:      userID,
		Categories:  categories,
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
		Message: "article succesfully updated",
	})

}

func (h *Handler) deleteArticle(ctx *gin.Context) {
	var id ID

	err := ctx.ShouldBindUri(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	userID := ctx.MustGet("userID").(int64)

	err = h.Srvs.DeleteArticle(ctx, userID, id.Value)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &ErrorResponse{
			http.StatusInternalServerError,
			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, &Response{
		Code:    http.StatusOK,
		Message: "article succesfully deleted",
	})
}
