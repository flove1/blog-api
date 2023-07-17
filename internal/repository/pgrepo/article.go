package pgrepo

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

func (p *Postgres) CreateArticle(ctx context.Context, a *entity.Article) error {
	articleQuery := fmt.Sprintf(`
		INSERT INTO %s (
			title,
			description,
			user_id
			)
		VALUES ($1, $2, $3)
		RETURNING id
		`, articlesTable)

	categoryQuery := fmt.Sprintf(`
		INSERT INTO %s (
			article_id,
			category_id
		) 
		VALUES (
			$1,
			(SELECT id FROM %s WHERE category_name = $2))
		`, articlesCategoriesTable, categoriesTable)

	tx, err := p.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	err = tx.QueryRow(ctx, articleQuery, a.Title, a.Description, a.UserID).Scan(&a.ID)
	if err != nil {
		return err
	}

	for _, category := range a.Categories {
		_, err := tx.Exec(ctx, categoryQuery, a.ID, strings.ToLower(category.Name))
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetArticleByID(ctx context.Context, id int64) (*entity.Article, error) {
	articleQuery := fmt.Sprintf(`
		SELECT 
			id,
			title,
			description,
			user_id
		FROM %s a
		WHERE a.id = $1
	`, articlesTable)

	categoriesQuery := fmt.Sprintf(`
		SELECT 
			c.id, 
			c.category_name 
		FROM %s ac 
		INNER JOIN %s c ON ac.category_id = c.id 
		WHERE ac.article_id = $1
	`, articlesCategoriesTable, categoriesTable)

	var article entity.Article
	err := p.Pool.QueryRow(ctx, articleQuery, id).Scan(&article.ID, &article.Title, &article.Description, &article.UserID)
	if err != nil {
		return nil, err
	}

	categories := make([]entity.Category, 0, 8)
	rows, err := p.Pool.Query(ctx, categoriesQuery, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category entity.Category
		rows.Scan(&category.ID, &category.Name)
		categories = append(categories, category)
	}

	article.Categories = categories
	return &article, nil
}

func (p *Postgres) GetAllArticles(ctx context.Context) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT 
			id,
			title,
			description,
			user_id,
			(SELECT ARRAY(
				SELECT 
					c.id 
				FROM %[1]s ac 
				INNER JOIN %[2]s c on ac.category_id = c.id 
				WHERE ac.article_id = a.id ORDER BY c.id) AS category_ids),
			(SELECT ARRAY(
				SELECT 
					c.category_name 
				FROM %[1]s ac 
				INNER JOIN %[2]s c on ac.category_id = c.id 
				WHERE ac.article_id = a.id ORDER BY c.id) AS category_names)
		FROM %[3]s a
	`, articlesCategoriesTable, categoriesTable, articlesTable)

	articles := make([]entity.Article, 0, 10)
	rows, err := p.Pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var article entity.Article
		caterogyIDs := make([]int64, 0, 5)
		caterogyNames := make([]string, 0, 5)

		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.UserID, &caterogyIDs, &caterogyNames)
		if err != nil {
			return nil, err
		}

		for index := 0; index < len(caterogyIDs); index++ {
			category := entity.Category{
				ID:   caterogyIDs[index],
				Name: caterogyNames[index],
			}
			article.Categories = append(article.Categories, category)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (p *Postgres) GetArticlesByUsername(ctx context.Context, username string) ([]entity.Article, error) {
	query := fmt.Sprintf(`
		SELECT 
			a.id,
			a.title,
			a.description,
			a.user_id,
			(SELECT ARRAY(
				SELECT 
					c.id 
				FROM %[1]s ac 
				INNER JOIN %[2]s c on ac.category_id = c.id 
				WHERE ac.article_id = a.id ORDER BY c.id) AS category_ids),
			(SELECT ARRAY(
				SELECT 
					c.category_name 
				FROM %[1]s ac 
				INNER JOIN %[2]s c on ac.category_id = c.id 
				WHERE ac.article_id = a.id ORDER BY c.id) AS category_names)
		FROM %[3]s a 
		INNER JOIN %[4]s u ON a.user_id = u.id
		WHERE u.username = $1
	`, articlesCategoriesTable, categoriesTable, articlesTable, usersTable)

	articles := make([]entity.Article, 0, 10)
	rows, err := p.Pool.Query(ctx, query, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var article entity.Article
		caterogyIDs := make([]int64, 0, 5)
		caterogyNames := make([]string, 0, 5)

		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.UserID, &caterogyIDs, &caterogyNames)
		if err != nil {
			return nil, err
		}

		for index := 0; index < len(caterogyIDs); index++ {
			category := entity.Category{
				ID:   caterogyIDs[index],
				Name: caterogyNames[index],
			}
			article.Categories = append(article.Categories, category)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (p *Postgres) UpdateArticle(ctx context.Context, a *entity.Article) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			title = COALESCE(NULLIF($2, ''), title),
			description = $3
		WHERE 
			id = $1 AND
			user_id = $4
	`, articlesTable)

	tag, err := p.Pool.Exec(ctx, query, a.ID, a.Title, a.Description, a.UserID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("article does not exists or does not belong to user")
	}

	return nil
}

func (p *Postgres) DeleteArticle(ctx context.Context, userID int64, articleID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = $1 and id = $2 
	`, articlesTable)

	tag, err := p.Pool.Exec(ctx, query, userID, articleID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("article does not exists or does not belong to user")
	}

	return nil
}
