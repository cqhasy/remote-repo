package main

import "fmt"

func max(s []int) int {
	var b int
	for i := 0; i < len(s)-1; i++ {
		if s[i] <= s[i+1] {
			b = s[i+1]
		} else {
			b = s[i]
		}
	}
	return b
}
func newmak(s []int) []int {
	b := make([]int, 0)
	for i := 0; i < len(s)-2; i++ {
		if s[i] <= s[i+1] {
			b = append(b, s[i])
		}
	}
	return b
}
func main() {
	var n, q int
	if _, err := fmt.Scan(&n, &q); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
	l := make([]int, n)
	for i := 0; i < n; i++ {
		if _, err := fmt.Scan(&l[i]); err != nil {
			fmt.Println(err)
			return
		}
	}
	for i := 0; i < q; i++ {
		var k int
		j := 1
		if _, err := fmt.Scan(&k); err != nil {
			fmt.Println(err)
			return
		}
		for {
			if k <= max(l) && len(l) > 0 {
				fmt.Println(j)
				break
			} else if len(l) < 0 {
				fmt.Println(-1)
			} else {
				k = k - max(l)
				l = newmak(l)
				j++
			}
		}

	}
}
