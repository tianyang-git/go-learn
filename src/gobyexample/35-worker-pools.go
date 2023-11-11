package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("before start len", len(jobs))
		fmt.Println("worker", id, "started  job", j, "len", len(jobs))
		time.Sleep(time.Second)
		fmt.Println()
		fmt.Println("worker", id, "finished job", j, "len", len(jobs))
		results <- j * 2
	}
}

func main() {
	start := time.Now() // 获取当前时间

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	fmt.Println("close jobs")
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	elapsed := time.Since(start)
	fmt.Println("该函数执行完成耗时：", elapsed)
}
