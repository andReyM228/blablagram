package main

import (
	"blablagram/logger/zaplog"
	"blablagram/repository"
	"blablagram/server"
	"blablagram/server/handlers"
	"blablagram/service"
	"context"
	"github.com/zeebo/errs"
	"net"
	"os"
)

func main() {
	log := zaplog.NewLog()
	ctx := context.Background()

	listener, err := net.Listen("tcp", ":7778")
	if err != nil {
		log.Fatal("creating a listener", err)
	}

	rep, err := repository.New(ctx, os.Getenv("MONGO_URL"))
	if err != nil {
		log.Fatal("creating a repository", err)
	}

	s, err := service.New(log, rep, os.Getenv("SALT"))
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
