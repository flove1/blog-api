package service

import (
	"context"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

func (m *Manager) CreateArticle(ctx context.Context, a *entity.Article) error {
	err := m.Repository.CreateArticle(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	article, err := m.Repository.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (m *Manager) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
	articles, err := m.Repository.GetAllArticles(ctx)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *Manager) GetArticlesByUsername(ctx context.Context, username string) ([]entity.Article, error) {
	articles, err := m.Repository.GetArticlesByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (m *Manager) UpdateArticle(ctx context.Context, a *entity.Article) error {
	err := m.Repository.UpdateArticle(ctx, a)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) DeleteArticle(ctx context.Context, userID int64, articleID int64) error {
	err := m.Repository.DeleteArticle(ctx, userID, articleID)
	if err != nil {
		return err
	}

	return nil
}
