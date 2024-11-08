package main

import (
	"fmt"
	"sync"
)

var mainWait sync.WaitGroup
var wait sync.WaitGroup

func main() {
	mainWait.Add(26)
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]int, 0, 26)
	for i := 0; i < 26; i++ {
		b = append(b, i)

	}
	for i := 0; i < 26; i++ {
		wait.Add(1)
		if i%2 == 0 && i >= 2 {
			go func() {
				fmt.Println(b[i-2])
				fmt.Println(b[i-1])
				wait.Done()
			}()

		} else {
			wait.Done()

		}
		wait.Wait()
		fmt.Printf("%c\n", a[i])

		mainWait.Done()
	}
	mainWait.Wait()
	fmt.Println(24)
	fmt.Println(25)
}
