package main

import "fmt"

func main() {
	cha1 := make(chan int,1)
	cha2 := make(chan int)

	/*cha1 <- 1
	cha2 <- 3*/

	select {
	case val1 := <-cha1:
		fmt.Println("cha1 val", val1)
	case cha2 <- 3:
		fmt.Println("put val to cha2")
	/*default:
		fmt.Println("default case")
	*/}


}
