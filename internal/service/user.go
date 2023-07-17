package service

import (
	"context"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
	"github.com/Tilvaldiyev/blog-api/pkg/jwtutil"
	"github.com/Tilvaldiyev/blog-api/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func (m *Manager) CreateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := m.Repository.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (m *Manager) CreateToken(ctx context.Context, u *entity.User, password string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := jwtutil.GenerateToken(u.ID, m.Config.AUTH.TokenExpiration, m.Config.AUTH.SigningKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (m *Manager) UpdateUser(ctx context.Context, u *entity.User) error {
	hashedPassword, err := util.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	err = m.Repository.UpdateUser(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) DeleteUser(ctx context.Context, userID int64) error {
	err := m.Repository.DeleteUser(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}
