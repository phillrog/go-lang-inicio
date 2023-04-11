package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	p := pessoa{"Teste", 30}
	fmt.Println(p)

	funcionario, erro := json.Marshal(p)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(funcionario)
	fmt.Println(bytes.NewBuffer(funcionario))

	c := map[string]string{
		"nome":  "teste 3",
		"idade": "99",
	}

	funcionario2, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(funcionario2)
	fmt.Println(bytes.NewBuffer(funcionario2))
}
