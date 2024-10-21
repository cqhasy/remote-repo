package main

import "fmt"

var a, b, i, v int

func mubiao(n []int, target int) []int {
	for i = 0; i < v; i++ {
		b = target - n[i]
		for j := i + 1; j < v; {
			if b == n[j] {
				return []int{i, j}
			}
			j++

		}
	}
	return nil
}
func main() {
	var t []int
	fmt.Scanf("%d%d", &a, &v)
	for i := 0; i < v; i++ {
		var l int
		fmt.Scanf("%d", &l)
		t = append(t, l)
	}
	fmt.Println(mubiao(t, a))
}

