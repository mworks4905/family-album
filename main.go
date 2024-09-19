package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/mworks4905/family-album/configs"
	"github.com/mworks4905/family-album/handlers"
	"github.com/mworks4905/family-album/s3"
	// "github.com/rs/zerolog/log"
)

type authed bool

var isLoggedIn authed = false

func main() {
	ctx := context.Background()
	configs.LoadEnvConfigs()

	// Start S3 client
	s3Client := s3.NewClient(ctx)

	// Create custom route handlers
	getPicturesHandler := handlers.GetPictures{S3: s3Client}
	getPictureHandler := handlers.GetPicture{S3: s3Client}

	// Create router and handles
	router := http.NewServeMux()
	router.Handle("GET /pictures", getPicturesHandler)
	router.Handle("GET /picture/{fileName}", getPictureHandler)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, isLoggedIn, true)
			return ctx
		},
	}

	fmt.Println("Listening on port: 9000")
	server.ListenAndServe()
}
