package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/mworks4905/family-album/configs"
	"github.com/mworks4905/family-album/internal/s3"
	"github.com/mworks4905/family-album/internal/server"
)

func main() {
	ctx := context.Background()
	configs.LoadEnvConfigs()

	// Start S3 client
	s3.NewClient(ctx, os.Getenv("AWS_BUCKET"))

	// Start server
	s := server.NewServer(":9000", ctx)
	fmt.Println("Listening on port: 9000")
	log.Fatal(s.ListenAndServe())
}
