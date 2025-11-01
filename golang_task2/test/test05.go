package test

import (
	"fmt"
	"sync"
	"time"
)

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

// 定义无参无返回值的函数类型为task
type Task func()

func TaskSchdle(tasks []Task) {
	var wg sync.WaitGroup

	for taskIndex, task := range tasks {
		wg.Add(1)
		go func(index int, taskParam Task) {
			start := time.Now().Unix()
			taskParam()
			end := time.Now().Unix()

			fmt.Printf("任务%v执行时间：%v", index, end-start)

			wg.Done()
		}(taskIndex, task)
	}

	wg.Wait()
}
