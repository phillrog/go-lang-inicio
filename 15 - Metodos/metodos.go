package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do usuário '%s' no banco de dados", u.nome)
}

func main() {
	usuario1 := usuario{"Teste", 20}

	fmt.Println(usuario1)
	usuario1.salvar()
}
