package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.NewRequest("POST", "http://muxithief.muxixyz.com/api/v1/login", nil)
	if err != nil {
		log.Fatal(err)
	}
	res.Header.Set("Code", "application/json")
	b := http.Client{}
	resp, err := b.Do(res)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	code := resp.Header.Get("Code")
	fmt.Println(code)
}
