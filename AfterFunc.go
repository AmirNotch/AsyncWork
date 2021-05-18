package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hi bratik")
}
func main() {
	timer := time.AfterFunc(5*time.Second, sayHello)

	fmt.Scanln()
	timer.Stop()

	fmt.Scanln()
}
