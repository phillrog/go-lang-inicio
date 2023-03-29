package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitgroup.Done()
	}()

	go func() {
		escrever("Programador GO")
		waitgroup.Done()
	}()

	waitgroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto, " - ", i)
		time.Sleep(time.Second)
	}
}
