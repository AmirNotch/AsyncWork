package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var totalOperation int32

func inc(){
	atomic.AddInt32(&totalOperation, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation", totalOperation)
}
