package main

import (
	"blablagram/config"
	"blablagram/logger/zaplog"
	"blablagram/repository"
	"blablagram/server"
	"blablagram/server/handlers"
	"blablagram/service"
	"context"
	"github.com/zeebo/errs"
	"net"
)

func main() {
	log := zaplog.NewLog()
	ctx := context.Background()

	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal("parsing config", err)
	}

	listener, err := net.Listen("tcp", ":7778")
	if err != nil {
		log.Fatal("creating a listener", err)
	}

	rep, err := repository.New(ctx, cfg.Mongo.Url)
	if err != nil {
		log.Fatal("creating a repository", err)
	}

	s, err := service.New(log, rep, cfg.Hash.Salt)
	if err != nil {
		log.Fatal("creating a service", err)
	}

	h := handlers.New(log, s)

	serv := server.NewServer(log, listener, h)

	runErr := serv.Run(ctx)
	closeErr := serv.Close()
	if runErr != nil || closeErr != nil {
		log.Error("server error", errs.Combine(runErr, closeErr))
	}
}
