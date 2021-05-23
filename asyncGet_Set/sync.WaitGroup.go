package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	goroutineNum = 7
	iterationNum = 5
)

func startWorkers(in int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := 0; j < iterationNum; j++ {
		fmt.Println(formatWorker(in, j))
		runtime.Gosched()
	}
}


func formatWorker(in, input int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), " ",
		strings.Repeat(" ", goroutineNum - in),
		"th", in,
		"recieved", input, strings.Repeat(" ", input))
}


func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < goroutineNum; i++ {
		wg.Add(1)
		go startWorkers(i, wg)
	}
	time.Sleep(time.Millisecond)

	wg.Wait()
}
