package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	iterations = 7
	gorountine = 5
)

func doSomeWork(in int){
	for j := 0; j < iterations; j++ {
		fmt.Printf(formatWork(in, j))
		runtime.Gosched()
	}
}

func main() {
	for i := 0; i < gorountine;i++ {
		go doSomeWork(i)
	}
	fmt.Scanln()
}
func formatWork(in,j int) string {
	return fmt.Sprintln(strings.Repeat(" ", in), "*",
		strings.Repeat(" ", gorountine-in),
		"th", in,
		"iter", j)
}


