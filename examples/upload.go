package main

import (
	"github.com/Yunshiro/gfile"
)

func main() {
	// wd, _ := os.Getwd()
	// fmt.Println("当前工作目录:", wd) // 检查程序运行位置
	gfile.Upload("examples/gopher.png")
}