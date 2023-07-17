package pgrepo

import (
	"context"
	"fmt"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
						username, -- 1 
						first_name, -- 2
						last_name, -- 3
						hashed_password -- 4
						)
		VALUES ($1, $2, $3, $4)
		`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.Username, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := fmt.Sprintf(`
			SELECT 
				id,
				username,
				first_name,
				last_name,
				hashed_password
			FROM %s
			WHERE 
				username = $1
			`, usersTable)

	user := &entity.User{}

	err := p.Pool.QueryRow(ctx, query, username).Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p *Postgres) UpdateUser(ctx context.Context, u *entity.User) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			first_name = COALESCE(NULLIF($2, ''), first_name),
			last_name = COALESCE(NULLIF($3, ''), last_name),
			hashed_password = COALESCE(NULLIF($4, ''), hashed_password)
		WHERE 
			id = $1
	`, usersTable)

	_, err := p.Pool.Exec(ctx, query, u.ID, u.FirstName, u.LastName, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) DeleteUser(ctx context.Context, userID int64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE id = $1
	`, usersTable)

	_, err := p.Pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	return nil
}
