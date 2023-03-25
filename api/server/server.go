package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dickykmrlh/protos/gen/go/user/v1/user_v1connect"
	"github.com/dickykmrlh/user/config"
	"github.com/dickykmrlh/user/database"
	service "github.com/dickykmrlh/user/internal/core/services"
	"github.com/dickykmrlh/user/internal/handler"
	"github.com/dickykmrlh/user/internal/repository"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	handler *handler.UserHandler
}

func New() (*Server, error) {
	userRepo := repository.NewUserRepo(database.GetDB())
	userSvc := service.NewUserService(userRepo)
	return &Server{
		handler: handler.NewUserHandler(userSvc),
	}, nil
}

func (s *Server) Run() error {
	httpCfg := config.GetConfig().GetHttpConfig()
	addr := fmt.Sprintf("%s:%d", httpCfg.Host, httpCfg.Port)
	mux := http.NewServeMux()
	path, handler := user_v1connect.NewUserServiceHandler(s.handler)
	mux.Handle(path, handler)
	log.Printf("listening at: %s", addr)
	return http.ListenAndServe(
		addr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
