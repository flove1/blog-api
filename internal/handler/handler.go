package handler

import (
	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/Tilvaldiyev/blog-api/internal/service"
)

type Handler struct {
	Srvs   service.Service
	Config *config.Config
}

func New(srvs service.Service, config *config.Config) *Handler {
	return &Handler{
		Srvs:   srvs,
		Config: config,
	}
}
