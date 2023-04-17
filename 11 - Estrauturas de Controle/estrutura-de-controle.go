package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	if outronumero := numero; outronumero > 0 {
		fmt.Println("Número maior que 0")
	} else if outronumero == 1 {
		fmt.Println("Numero igual a 1")
	} else {
		fmt.Println("Menor que 0")
	}
}
