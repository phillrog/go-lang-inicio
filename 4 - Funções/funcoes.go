package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func subtrair(n1 int8, n2 int8) int8 {
	return n2 - n1
}

func calculosmatematicos(n1, n2 int8) (int8, int8) {
	return somar(n1, n2), subtrair(n2, n1)
}

func main() {
	fmt.Println(somar(10, 100))
	fmt.Println(calculosmatematicos(100, 11))
	resultado1, _ := calculosmatematicos(99, 1)
	fmt.Println(resultado1)
	_, resultado2 := calculosmatematicos(123, 100)
	fmt.Println(resultado2)
}
