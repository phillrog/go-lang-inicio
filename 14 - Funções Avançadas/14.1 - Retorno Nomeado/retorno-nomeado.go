package main

import "fmt"

func calculosmatematicos(n1, n2 int8) (soma int8, subtracao int8) {
	soma = n1 + n2
	subtracao = n2 - n1
	return soma, subtracao
}

func main() {
	fmt.Println("Retorno Nomeado")

	soma, subtracao := calculosmatematicos(10, 100)
	fmt.Println(soma, subtracao)
}
