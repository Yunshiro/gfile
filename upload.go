package gfile

import (
	"fmt"
)

// MultipartForm represents one part of file parts in HTTP request.
type MultipartForm struct {
	ContentType 		string
	Content 			any
	Boundary 			string
	ContentDisposition 	string
}

func Upload() {
	fmt.Println("upload success")
	
}