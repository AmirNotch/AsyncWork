package main

import "fmt"

func main(){
	cha1 := make(chan int,2)

	go func(in chan int) {
		val := <- in
		fmt.Println("GO: get some info", val)
		fmt.Println("GO: get some info")
	}(cha1)

	cha1 <- 42
	cha1 <- 1005006600

	fmt.Println("MAIN: after put to chan")
	fmt.Scanln()
}
