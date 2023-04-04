package main

import (
	"fmt"
	"time"
)

func main() {
	c1, c2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Canal 2"
		}
	}()

	for {
		select {
		case m1 := <-c1:
			{
				fmt.Println(m1)
			}
		case m2 := <-c2:
			{
				fmt.Println(m2)
			}
		}
	}

}
