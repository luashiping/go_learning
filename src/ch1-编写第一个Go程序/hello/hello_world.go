package main

import (
	"fmt"
	"os"
)

// Go中main函数不支持任何返回值
// 通过os.Exit来返回状态
// main函数不支持传入参数
// 在程序中直接通过os.Args获取命令行参数
func main() {
	// fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
	fmt.Println("Hello World")
	// os.Exit(-1)
}
