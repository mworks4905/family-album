package s3

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Client struct {
	Client *s3.Client
	Bucket string
}

func NewClient(ctx context.Context) *S3Client {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	myS3 := S3Client{
		Client: s3.NewFromConfig(cfg),
		Bucket: os.Getenv("AWS_BUCKET"),
	}

	return &myS3
}

func (s *S3Client) List(prefix string) *s3.ListObjectsV2Output {
	input := s3.ListObjectsV2Input{
		Bucket: &s.Bucket,
	}

	if prefix != "" {
		input.Prefix = &prefix
	}

	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := s.Client.ListObjectsV2(context.TODO(), &input)
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func (s *S3Client) Read(key string) *s3.GetObjectOutput {
	input := s3.GetObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
	}

	output, err := s.Client.GetObject(context.TODO(), &input)
	if err != nil {
		// log.Fatal(err)
		fmt.Printf("Error finding file: %v", err)
	}

	return output
}

func (s *S3Client) Upload(key string, data io.Reader) *s3.PutObjectOutput {
	input := s3.PutObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
		Body:   data,
	}

	output, err := s.Client.PutObject(context.TODO(), &input)
	if err != nil {
		fmt.Printf("Error uploading data: %v\n", data)
	}

	return output
}

func (s *S3Client) MultipartUpload(key string, data io.Reader) {
	// do multipart upload
}
