package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := errors.New("file not exist")
	if err != nil {
		fmt.Println(err)
	}

	_, errOpen := os.Open("os3.go")
	if errOpen != nil {
		fmt.Println(os.IsNotExist(errOpen))
		fmt.Println(errOpen)
	}

	var path string
	if os.IsPathSeparator('/') {
		path = "/"
	} else {
		path = "//"
	}
	fmt.Println(path)
	dir, _ := os.Getwd()
	fmt.Println(dir)
	// errMake := os.Mkdir(dir+path+"mk", os.ModePerm)
	// if errMake != nil {
	// 	fmt.Println(errMake)
	// 	return
	// }
	// fmt.Println("创建目录" + dir + path + "mk成功")

	errMake := os.MkdirAll(dir+path+"A/B", os.ModePerm)
	if errMake != nil {
		fmt.Println(errMake)
		return
	}
	fmt.Println("创建文件夹")
}
