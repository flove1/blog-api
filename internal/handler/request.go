package handler

type ID struct {
	Value int64 `uri:"id" binding:"required"`
}

type Username struct {
	Value string `uri:"username" binding:"required,min=5,max=50"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserRequest struct {
	Username  string `json:"username" binding:"required,min=5,max=50"`
	Password  string `json:"password" binding:"required,min=4,max=32"`
	FirstName string `json:"first_name" binding:"required,min=1,max=50"`
	LastName  string `json:"last_name" binding:"required,min=1,max=50"`
}

type UpdateUserRequest struct {
	Password  string `json:"password" binding:"min=4,max=8"`
	FirstName string `json:"first_name" binding:"min=1,max=50"`
	LastName  string `json:"last_name" binding:"min=1,max=50"`
}

type CreateArticleRequest struct {
	Title       string   `json:"title" binding:"required,min=5,max=100"`
	Description string   `json:"description" binding:"required"`
	Categories  []string `json:"categories" binding:"required"`
}

type UpdateArticleRequest struct {
	ID
	Title       string   `json:"title" binding:"max=100"`
	Description string   `json:"description"`
	Categories  []string `json:"categories"`
}
