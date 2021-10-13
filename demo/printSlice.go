package demo

import "fmt"

func PrintSlice() {
	arrayList := make([]int, 0, 5)
	c := cap(arrayList)
	fmt.Printf("当前容量：%d\n", c)
	fmt.Printf("当前地址：%p\n", arrayList)
	for i := 0; i < 10; i++ {
		arrayList = append(arrayList, i)
	}
	//如果容量改变了，切片会扩容，实际上是复制了原始值的新数组，所以数组的地址会改变
	c = cap(arrayList)
	fmt.Printf("当前容量：%d\n", c)
	fmt.Printf("当前地址：%p\n", arrayList)

	scores := []int{1, 2, 3, 4, 5}
	scores = scores[0:2] //[0:2)左闭右开
	fmt.Println(scores)
}
