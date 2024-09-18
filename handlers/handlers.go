package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/mworks4905/family-album/s3"
)

func GetPictures(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /pictures")

	var S3 = r.Context().Value("S3").(*s3.S3Client)

	contents := S3.List("")
	for _, object := range contents.Contents {
		fmt.Println(*object.Key)
	}
}

func GetPicture(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /picture/{fileName}")
	fileName := r.PathValue("fileName")
	var S3 = r.Context().Value("S3").(*s3.S3Client)

	res := S3.Read(fileName)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	contents := buf.String()

	fmt.Println(contents)
}
