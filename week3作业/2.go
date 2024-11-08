package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type m struct {
	number int
	id     int
}

var a chan m
var waitgroup sync.WaitGroup

func main() {
	waitgroup.Add(20)
	a = make(chan m, 20)
	defer close(a)
	for i := 1; i < 21; i++ {
		go func() {
			randomNumber := rand.Intn(100)
			a <- m{
				number: randomNumber,
				id:     i}
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()
	fmt.Println("排序前：")
	d := make([]m, 0, 20)
	for i := 0; i < 20; i++ {
		h := <-a
		d = append(d, h)
		fmt.Println(h.number, h.id)
	}
	fmt.Println("排序后：")
	for i := 0; i < 20; i++ {
		if i != d[i].id-1 {
			c := d[i]
			d[i] = d[d[i].id-1]
			d[c.id-1] = c
		} else {
			continue
		}
	}
	for _, v := range d {
		fmt.Println(v.number, v.id)
	}
}
