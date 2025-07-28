package gfile

import (
	"fmt"
	"log"
	"os"
)

// MultipartForm represents one part of file parts in HTTP request.
type MultipartForm struct {
	ContentType 		string
	Content 			any
	Boundary 			string
	ContentDisposition 	string
}

func Transfer(filename string) []byte {
	fileBytes := toBytes(filename)
	return fileBytes
}

func toBytes(filename string) []byte {
	bts, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("file to bytes error:%v", err)
	}
	return bts
}



func Upload(filename string) {
	fmt.Println("upload success")
	fmt.Printf("bytes: %b\n", Transfer(filename))
}