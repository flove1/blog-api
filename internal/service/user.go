package service

import (
	"context"

	"github.com/Tilvaldiyev/blog-api/internal/entity"
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

func (m *Manager) Login(ctx context.Context, username, password string) (*entity.User, error) {
	user, err := m.Repository.Login(ctx, username, password)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
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
