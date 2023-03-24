package server

import (
	"github.com/dickykmrlh/user/internal/handler"
)

type Server struct {
	handler *handler.UserHandler
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	/*
		mux := http.NewServeMux()
		path, handler := user_v1connect.NewUserServiceHandler(s.handler)
		mux.Handle(path, handler)
		return http.ListenAndServe(
			address,
			// Use h2c so we can serve HTTP/2 without TLS.
			h2c.NewHandler(mux, &http2.Server{}),
		)
	*/
	return nil
}
