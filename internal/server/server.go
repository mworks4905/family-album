package server

import (
	"context"
	"net"
	"net/http"

	"github.com/mworks4905/family-album/internal/handlers"
)

type authed bool

var isLoggedIn authed = false

func NewServer(addr string, ctx context.Context) *http.Server {
	r := http.NewServeMux()

	r.HandleFunc("GET /pictures", handlers.GetPictures)
	r.HandleFunc("GET /picture", handlers.GetPicture)

	return &http.Server{
		Addr:    addr,
		Handler: r,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, isLoggedIn, true)
			return ctx
		},
	}
}
