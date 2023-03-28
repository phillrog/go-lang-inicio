package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops")

	i := 0

	for i < 10 {
		i++
		fmt.Println(i)
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	for j := 0; j < 10; j += 2 {
		fmt.Println(j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"joão", "paulo", "ana"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, nome := range "PALAVRA" {
		fmt.Println(indice, string(nome))
	}

	usuario := map[string]string{
		"nome":      "TESTE",
		"sobrenome": "TESTE2",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
