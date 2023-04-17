package main

import "fmt"

func soma(n ...int) int {
	total := 0
	for _, numero := range n {
		total += numero
	}

	return int(total)
}

func escrever(msg string, n ...int) {
	for _, numero := range n {
		fmt.Println(msg, numero)
	}
}

func main() {
	totalSoma := soma(10, 20, 30, 40, 50, 60, 70, 80, 90, 100)

	fmt.Println(totalSoma)

	escrever("Olá Mundo! ", 10, 20, 30, 40, 50, 60, 70, 80, 90, 100)
}
