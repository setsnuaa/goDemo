package demo

import "fmt"

func PrintDefer() {
	fmt.Println("程序开始执行")
	defer fmt.Println("程序结束后执行该函数")
	fmt.Println("程序执行结束")
}
