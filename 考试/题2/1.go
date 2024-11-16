package main

import "fmt"

func jue(a, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
func least(a, b int) int {
	if a >= b {
		if jue(a, b) == 2 || jue(a, b) == 1 || jue(a, b) == 0 {
			return 0
		} else {
			return a - b - 2

		}

	} else {
		return jue(a, b)
	}

}
func main() {
	var n int
	slice := make([]int, 0)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var t int
		fmt.Scanf("%d", &t)
		slice = append(slice, t)
	}
	var sum = 0
	var fin = 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			b := least(slice[i], slice[j])
			sum += b
		}
		fin += sum

	}
	fmt.Println(fin)
}
