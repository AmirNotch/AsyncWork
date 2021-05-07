package main

import "fmt"

func main() {
	cha1 := make(chan int,2)
	cha1 <- 1
	cha1 <- 2
	cha2 := make(chan int,2)
	cha2 <- 3
LOOP:
	for {
		select {
		case val1 := <-cha1:
			fmt.Println("Channel 1 ", val1)
		case val2 := <-cha2:
			fmt.Println("Channel 2 ", val2)
		default:
			break LOOP
		}
	}

}
