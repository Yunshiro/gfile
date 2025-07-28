package gfile

import (
	"fmt"
	"os"
)

func toFile(bts []byte, name string) bool{
	err := os.WriteFile(name, bts, 0644)
	fmt.Println("err: ", err)
	return err == nil
}