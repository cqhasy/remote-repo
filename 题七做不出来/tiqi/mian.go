package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 发送 GET 请求
	resp, err := http.Get("http://121.43.151.190:8000/paper")
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close() // 确保在函数结束时关闭响应体

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// 打印响应体内容
	fmt.Println(string(body))
}
