package pgrepo

import (
	"context"
	"fmt"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

func (p *Postgres) GetCategories(ctx context.Context) ([]entity.Category, error) {
	query := fmt.Sprintf(`
		SELECT
			id,
			category_name
		FROM %s
	`, categoriesTable)

	categories := make([]entity.Category, 0, 10)
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category entity.Category
		err = rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
		if err != nil {
			return nil, err
		}
	}

	return categories, nil
}
