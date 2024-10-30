package main

import "fmt"

func jishu() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}
func main() {
	a := jishu()
	fmt.Println(a())
	fmt.Println(a())
}
