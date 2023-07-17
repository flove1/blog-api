package service

import (
	"context"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

type Service interface {
	CreateUser(ctx context.Context, u *entity.User) error
	Login(ctx context.Context, username, password string) (*entity.User, error)
	UpdateUser(ctx context.Context, u *entity.User) error
	DeleteUser(ctx context.Context, id int64) error

	CreateArticle(ctx context.Context, a *entity.Article) error
	UpdateArticle(ctx context.Context, a *entity.Article) error
	DeleteArticle(ctx context.Context, userID int64, articleID int64) error
	GetArticleByID(ctx context.Context, id int64) (*entity.Article, error)
	GetAllArticles(ctx context.Context) ([]entity.Article, error)
	GetArticlesByUsername(ctx context.Context, username string) ([]entity.Article, error)

	GetCategories(ctx context.Context) ([]entity.Category, error)
}
