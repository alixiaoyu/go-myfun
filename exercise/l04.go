package main

import "fmt"

func main() {
	/*字符串转为字节数组*/
	var username string = "dequan,你好"
	for _, v := range username {
		fmt.Println("---->", v)

	}
	nameCharAr := []rune(username) //把字符串转为字节数组，每一位存储的是该字符对应的ASCII码
	//针对英文，一个字符占用1个字节；针对汉字，utf8的，一个字符占用3个字节
	// var len = len(nameCharAr)

	// for i := 0; i < len; i++ {
	// 	fmt.Println(nameCharAr[i], "===>", string(nameCharAr[i]))
	// }
	for i, v := range nameCharAr {
		fmt.Println("---->", v)
		fmt.Println(nameCharAr[i], "===>", string(nameCharAr[i]))
	}
}
