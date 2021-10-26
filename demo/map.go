package demo

import (
	"fmt"
	"strconv"
)

func PrintMap() {
	myMap := make(map[string]int)
	myMap["tony"] = 30
	for key, value := range myMap {
		//迭代映射是没有顺序。每次迭代查找将会随机返回键值对。
		fmt.Println("name:" + key + " age:" + strconv.Itoa(value))
	}
}
