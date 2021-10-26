package demo

import (
	"fmt"
	"reflect"
)

func PrintType(a interface{}) {
	fmt.Println("the type of the variable is", reflect.TypeOf(a))
}
