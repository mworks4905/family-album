package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/mworks4905/family-album/internal/s3"
)

type IS3 interface {
	List(prefix string) *awsS3.ListObjectsV2Output
	Read(key string) *awsS3.GetObjectOutput
}

var S3 IS3 = s3.Client

func GetPictures(w http.ResponseWriter, r *http.Request) {
	contents := s3.Client.List("")
	var results []string
	for _, object := range contents.Contents {
		fmt.Println(*object.Key)
		results = append(results, *object.Key)
	}
	res := struct {
		Contents []string
	}{
		Contents: results,
	}
	json.NewEncoder(w).Encode(&res)
}

func GetPicture(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	results := s3.Client.Read(key)

	buf := new(bytes.Buffer)
	buf.ReadFrom(results.Body)
	contents := buf.String()

	// w.Write([]byte(contents))
	res := struct {
		Contents string
	}{
		Contents: contents,
	}
	json.NewEncoder(w).Encode(&res)
}
