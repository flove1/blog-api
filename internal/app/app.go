package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/Tilvaldiyev/blog-api/internal/config"
	"github.com/Tilvaldiyev/blog-api/internal/handler"
	"github.com/Tilvaldiyev/blog-api/internal/repository/pgrepo"
	"github.com/Tilvaldiyev/blog-api/internal/service"
	"github.com/Tilvaldiyev/blog-api/pkg/httpserver"
)

func Run(cfg *config.Config) error {
	db, err := pgrepo.New(
		pgrepo.WithHost(cfg.DB.Host),
		pgrepo.WithPort(cfg.DB.Port),
		pgrepo.WithDBName(cfg.DB.DBName),
		pgrepo.WithUsername(cfg.DB.Username),
		pgrepo.WithPassword(cfg.DB.Password),
	)
	if err != nil {
		log.Printf("connection to DB err: %s", err.Error())
		return err
	}
	log.Println("connection success")

	srvs := service.New(db, cfg)
	hndlr := handler.New(srvs, cfg)
	server := httpserver.New(
		hndlr.InitRouter(),
		httpserver.WithPort(cfg.HTTP.Port),
		httpserver.WithReadTimeout(cfg.HTTP.ReadTimeout),
		httpserver.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		httpserver.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout),
	)

	server.Start()
	log.Println("server started")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case s := <-interrupt:
		log.Printf("signal received: %s", s.String())
	case err = <-server.Notify():
		log.Printf("server notify: %s", err.Error())
	}

	err = server.Shutdown()
	if err != nil {
		log.Printf("server shutdown err: %s", err)
	}

	return nil
}
