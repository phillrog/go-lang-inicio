package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}
type estudande struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Paulo", "Henrique", 36, 176}

	fmt.Println(p1)

	e1 := estudande{p1, "Técnico Informática", "Centro Paula Souza"}
	fmt.Println(e1)
}
