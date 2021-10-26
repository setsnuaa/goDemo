package demo

import "fmt"

// PrintMixCharacter 中英文混合字符串，统计字数
// 中文字符串使用len(str)会显示unicode编码大小，一个中文=3字节
func PrintMixCharacter(s string) {
	mixStr := []rune(s)
	fmt.Printf("\"%s\"的长度为：%d", string(mixStr), len(mixStr))
}
