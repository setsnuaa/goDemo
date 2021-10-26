package demo

import "fmt"

type MyInterface interface {
	Print(msg string)
}

type MyInterfaceImpl struct {
}

func (i *MyInterfaceImpl) Print(msg string) {
	fmt.Println("接口实现：" + msg)
}
