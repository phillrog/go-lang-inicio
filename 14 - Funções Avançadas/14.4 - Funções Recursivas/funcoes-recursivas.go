package main

import "fmt"

func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}

	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	fmt.Println("Funções Recursivas")

	posicao := uint(30)

	for i := uint(0); i < posicao; i++ {
		fmt.Println(fibonacci(i))
	}
}
