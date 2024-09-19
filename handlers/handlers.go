package handlers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/mworks4905/family-album/s3"
)

type Picture struct {
	Title string `json:"title"`
	Date  string `json:"data"`
}

type GetPictures struct {
	S3 *s3.S3Client
}

func (gps GetPictures) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /pictures")

	contents := gps.S3.List("")
	for _, object := range contents.Contents {
		fmt.Println(*object.Key)
	}
}

type GetPicture struct {
	S3 *s3.S3Client
}

func (gp GetPicture) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /picture/{fileName}")
	fileName := r.PathValue("fileName")

	res := gp.S3.Read(fileName)

	buf := new(bytes.Buffer)
	buf.ReadFrom(res.Body)
	contents := buf.String()

	fmt.Println(contents)
}

// func UploadPicture(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("POST /picture")
// 	picture := Picture{}
// 	err := json.NewDecoder(r.Body).Decode(&picture)
// 	if err != nil {
// 		fmt.Printf("there was an error: %v\n", err)
// 		return
// 	}

// 	fmt.Println(picture)
// 	var S3 = r.Context().Value("S3").(*s3.S3Client)
// 	res = S3.Upload()
// }
