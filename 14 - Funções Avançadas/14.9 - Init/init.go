package main

import "fmt"

var n int

func main() {
	fmt.Println("Função main")
	n = 20

	fmt.Println(n)
}

func init() {
	fmt.Println("Função init")

	n = 20000

	fmt.Println(n)
}
