package main

import (
	"fmt"
	"sync"
	"time"
)

func shishi(a *int, b int) int {
	return *a - b
}
func main() {
	var M = 2000
	var n = 10
	var k = 200
	var ch = make(chan int)
	defer close(ch)
	var wait sync.WaitGroup
	wait.Add(2000)
	var mu sync.Mutex
	for i := 0; i < n; i++ {
		go func() {
			mu.Lock()
			M = shishi(&M, k)
			ch <- M
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	go func() {
		for M != 0 {
			fmt.Println(<-ch)
			wait.Done()
		}
	}()
	wait.Wait()
	
}
