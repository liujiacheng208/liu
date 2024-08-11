package main

import (
	"fmt"
	"reflect"
)

func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	fmt.Println(reType)
	reValut := reflect.ValueOf(i)
	fmt.Println(reValut)
}0
func main() {
	var num int = 100
	testReflect(num)

}
