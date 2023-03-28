package main

import "fmt"

func main() {
	// ARITIMÉTICOS
	soma := 1 + 2
	subtracao := 2 - 1
	divisao := 10 / 2
	multiplicacao := 30 * 999
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// ATRIBUIÇÃO

	var a string = "a"
	a2 := "a2"

	fmt.Println(a, a2)

	// RELACIONAIS

	fmt.Print(1 > 2)
	fmt.Print(1 < 2)
	fmt.Print(1 >= 2)
	fmt.Print(1 <= 2)
	fmt.Print(1 == 2)
	fmt.Print(1 != 2)

	// LÓGICOS
	var verdadeiro, falso bool = true, false
	fmt.Println(verdadeiro || falso)
	fmt.Println(verdadeiro && falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// UNÁRIOS
	n1 := 0
	n1++
	fmt.Println(n1)
	n1 += n1
	n1--
	fmt.Println(n1)
	n1 -= n1
	fmt.Println(n1)

	if 1 > 2 {
		fmt.Println("sim")
	} else {
		fmt.Println("não")
	}

}
