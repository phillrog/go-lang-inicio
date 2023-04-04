package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo"
	canal <- "Programando com GO"

	m1 := <-canal
	m2 := <-canal
	fmt.Println(m1)
	fmt.Println(m2)
}
