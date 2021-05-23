package main

import (
	"fmt"
	"time"
)

func getComments() chan string {

	result := make(chan string, 1)
	go func(out chan <-string) {
		time.Sleep(2 * time.Second)
		fmt.Println()
		out <- "32 комментария"
	}(result)
	return result
}

func getPage() {
	var resultCh chan string = getComments()

	time.Sleep(1 * time.Second)
	fmt.Println("get related articles")

	return

	commentsData := <- resultCh
	fmt.Println("main goroutine:",commentsData)
}

func main() {
	getPage()
}

