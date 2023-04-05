package main

import (
	"fmt"
	"go-lang-inicio/19-Testes-automatizados/1-Introducao/enderecos"
)

func main() {
	endereco := enderecos.TipoEndereco("Avenida Presidente Vargas")

	fmt.Println(endereco)
}
