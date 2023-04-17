package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcção closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"

	fmt.Println(texto)

	funcaoNova := closure()

	funcaoNova()
}
