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

type S3Client string

var s3Client S3Client = "S3"

func main() {
	configs.LoadEnvConfigs()
	S3 := s3.InitClient()

	router := http.NewServeMux()

	router.HandleFunc("GET /pictures", handlers.GetPictures)
	router.HandleFunc("GET /picture/{id}", handlers.GetPicture)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.Background()
			ctx = context.WithValue(ctx, s3Client, S3)
			return ctx
		},
	}

	fmt.Println("Listening on port: 9000")
	server.ListenAndServe()
}
