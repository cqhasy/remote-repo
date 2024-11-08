package main

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

var lock sync.Mutex
var lock2 sync.Mutex
var wait sync.WaitGroup

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)
	defer close(msgs)

	// 这里模拟不断生产信息
	go func() {
		for i := 0; ; i++ {
			lock.Lock()
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}
			// 每次发送信息会暂停 0.01 秒以模拟真实场景
			time.Sleep(100 * time.Millisecond)
			lock.Unlock()
		}
	}()

	// 不断接收消息进行消费
	for msg := range msgs {
		// 添加新的值到 consumeMSG 中
		consumeMSG = append(consumeMSG, msg)

		// 如果消息数量达到批量消费的要求
		if len(consumeMSG) >= ConsumeNum {
			//h := len(consumeMSG)
			wait.Add(1)
			// 异步消费
			go func() {
				lock2.Lock()
				m := consumeMSG[:ConsumeNum] // 获取前 ConsumeNum 条消息
				fn(m)
				wait.Done()

				lock2.Unlock()

			}()
			// 更新上次消费时间
			lastConsumeTime = time.Now()
			wait.Wait()
			// 清空已消费的数据
			consumeMSG = consumeMSG[ConsumeNum:]

		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			// 如果上次消费已经超过 5 分钟且有未处理的消息
			if len(consumeMSG) > 0 {
				// 异步消费
				go func() {
					lock2.Lock()
					m := consumeMSG[:] // 获取所有待处理消息
					fn(m)
					lock2.Unlock()
				}()
				// 更新上次消费时间
				lastConsumeTime = time.Now()
				// 清空数据
				consumeMSG = nil
			}
		}
	}
}

func fn(m []MSG) {
	fmt.Printf("本次消费了 %d 条消息\n", len(m))
}
