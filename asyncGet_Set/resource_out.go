package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	goroutineNums = 7
	iterationNums = 5
	quotaLimit    = 2
)

func startWorking(in int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{}
	defer wg.Done()
	for j := 0; j < iterationNums; j++ {
		fmt.Printf(formatWorking(in, j))
		if j % 2 == 0 {
			<- quotaCh
			quotaCh <- struct{}{}
		}
		runtime.Gosched()
	}
	<-quotaCh
}

func main() {
	wg := &sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit)
	for i := 0; i < goroutineNums; i++ {
		wg.Add(1)
		go startWorking(i, wg, quotaCh)
	}
	time.Sleep(time.Millisecond)
	wg.Wait()

}
func formatWorking(in, input int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), " ",
		strings.Repeat(" ", goroutineNums-in),
		"th", in,
		"recieved", input, strings.Repeat(" ", input))
}
