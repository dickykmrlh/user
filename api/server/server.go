package server

import (
	"fmt"
	"net/http"

	"github.com/dickykmrlh/protos/gen/go/user/v1/user_v1connect"
)

type Server struct {
}

func Run() {
	mux := http.NewServeMux()
	path, handler := user_v1connect.NewUserServiceHandler(&petStoreServiceServer{})
	mux.Handle(path, handler)
	fmt.Println("... Listening on", address)
	http.ListenAndServe(
		address,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
