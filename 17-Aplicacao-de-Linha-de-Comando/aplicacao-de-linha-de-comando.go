package main

import (
	"fmt"
	"go-lang-inicio/17-Aplicacao-de-Linha-de-Comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()
	fmt.Println(aplicacao)
	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
