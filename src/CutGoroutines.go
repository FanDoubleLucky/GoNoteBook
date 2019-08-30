package main

import "fmt"

func main() {

	PV := make(chan int)
	go func() {
		for i := 1; i < 10; i++ {
			fmt.Println(i * 100)
		}

		PV <- 1
	}()

	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	<-PV
}
