package main

import (
	"fmt"
	"time"
)

//func double

func main() {
	x := 5
	y := func() int {
		return x * 2
	}()
	fmt.Println(y)

	c := make(chan int)

	go func() {
		fmt.Println("Starting...")
		time.Sleep(time.Second * 5)
		fmt.Println("End")
		c <- 1
	}()
	<-c
}
