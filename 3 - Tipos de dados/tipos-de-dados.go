package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numero int64 = 100000000
	fmt.Println(numero)

	var numero2 uint = 100000
	fmt.Println(numero2)

	var numero3 rune = 123456
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numero5 = strconv.FormatInt(int64(123), 10)
	fmt.Println(numero5)

	var numero6 = strconv.Itoa(123)
	fmt.Println(numero6)

	var numero7, err = strconv.Atoi("44")

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	fmt.Println(numero7)
}
