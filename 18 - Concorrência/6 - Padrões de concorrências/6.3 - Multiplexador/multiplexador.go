package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplexador(escrever("Olá Mundo!"), escrever("Programando em GO!"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(c1, c2 <-chan string) <-chan string {
	c3 := make(chan string)

	go func() {
		for {
			select {
			case msg := <-c1:
				{
					c3 <- msg
				}
			case msg := <-c2:
				{
					c3 <- msg
				}
			}
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return c3
}

func escrever(texto string) <-chan string {
	c1 := make(chan string)
	go func() {
		for {
			c1 <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return c1
}
