package main

import "fmt"

type endereco struct {
	logradour string
	numero    string
}

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "teste"
	u.idade = 33
	fmt.Println(u)

	u2 := usuario{"Teste 2", 45, endereco{"Rua Pestalozzi", "598"}}
	fmt.Println(u2)

	u3 := usuario{nome: "Teste 3"}
	fmt.Println(u3)
}
