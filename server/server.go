package server

import (
	"blablagram/logger"
	logger2 "blablagram/logger"
	"blablagram/server/handlers"
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/zeebo/errs"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

var (
	// Error is an error class that indicates internal http server error.
	Error = errs.Class("web server error")
)

// Server represents console web server.
type Server struct {
	log logger2.Logger

	listener net.Listener
	server   http.Server
	handlers *handlers.Handlers
}

// NewServer is a constructor for console web server.
func NewServer(log logger.Logger, listener net.Listener, handlers *handlers.Handlers) *Server {
	server := &Server{
		log:      log,
		listener: listener,
	}

	router := mux.NewRouter()
	router.HandleFunc("/status", handlers.Status).Methods(http.MethodPost)

	router.HandleFunc("/reg", handlers.RegisterUser).Methods(http.MethodPost)
	router.HandleFunc("/login", handlers.LoginUser).Methods(http.MethodPost)

	router.HandleFunc("/create_post", handlers.CreatePost).Methods(http.MethodPost)
	router.HandleFunc("/delete_post", handlers.DeletePost).Methods(http.MethodPost)
	router.HandleFunc("/update_post", handlers.UpdatePost).Methods(http.MethodPost)
	router.HandleFunc("/get_post", handlers.GetPost).Methods(http.MethodPost)
	router.HandleFunc("/get_posts", handlers.GetPosts).Methods(http.MethodPost)

	server.server = http.Server{
		Handler:     cors.Default().Handler(router),
		ReadTimeout: time.Hour,
		IdleTimeout: time.Hour,
	}

	return server
}

// Run starts the server that host api endpoint.
func (server *Server) Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return Error.Wrap(server.server.Shutdown(context.Background()))
	})
	group.Go(func() error {
		defer cancel()
		err := server.server.Serve(server.listener)
		isCancelled := errs.IsFunc(err, func(err error) bool { return errors.Is(err, context.Canceled) })
		if isCancelled || errors.Is(err, http.ErrServerClosed) {
			err = nil
		}
		return Error.Wrap(err)
	})

	return Error.Wrap(group.Wait())
}

// Close closes server and underlying listener.
func (server *Server) Close() error {
	return Error.Wrap(server.server.Close())
}
