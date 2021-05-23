package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const goroutinesNum = 3

func startWorker(workerNum int, in <-chan string){
	for input := range in {
		fmt.Printf(formatWork(workerNum, input))
		runtime.Gosched()
	}
	printFinishWork(workerNum)
}

func formatWork(in int, input string) string {
	return fmt.Sprintln(strings.Repeat("  ", in), " ",
		strings.Repeat(" ", goroutinesNum - in),
		"th", in,
		"recieved", input)
}

func printFinishWork(in int ) {
	fmt.Println(strings.Repeat("  ", in), " ",
		strings.Repeat(" ", goroutinesNum - in),
		"===", in,
		"finished")
}

func main() {
	workerInput := make(chan string, 2)
	for i := 0; i < goroutinesNum;i++ {
		go startWorker(i, workerInput)
	}

	months := []string{
		"Январь","Февраль","Март",
		"Апрель","Май","Июнь",
		"Июль","Август","Сентябрь",
		"Октябрь","Ноябрь","Декабрь",
	}

	for _,monthName := range months {
		workerInput <- monthName
	}
	close(workerInput)
	time.Sleep(time.Millisecond)
}
