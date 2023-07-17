package handler

import "github.com/gin-gonic/gin"

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	userV1 := v1.Group("/user")
	articleV1 := v1.Group("/article")
	categoryV1 := v1.Group("/category")

	auth := authenticationMiddleware(h.Config)

	userV1.POST("/register", h.createUser)
	userV1.POST("/login", h.login)

	userV1.Use(auth).PATCH("/update", h.updateUser)
	userV1.Use(auth).DELETE("/delete", h.deleteUser)

	articleV1.GET("", h.getAllArticles)
	articleV1.GET("/:id", h.getArticleById)
	articleV1.GET("/by/:username", h.getArticlesByUsername)

	articleV1.Use(auth).POST("", h.createArticle)
	articleV1.Use(auth).PATCH("/:id", h.updateArticle)
	articleV1.Use(auth).DELETE("/:id", h.deleteArticle)

	categoryV1.GET("", h.getCategories)

	return router
}
