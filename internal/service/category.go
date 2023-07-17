package service

import (
	"context"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

func (m *Manager) GetCategories(ctx context.Context) ([]entity.Category, error) {
	categories, err := m.Repository.GetCategories(ctx)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
