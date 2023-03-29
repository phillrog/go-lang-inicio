package main

import (
	"go-lang-inicio/17-Aplicacao-de-Linha-de-Comando/app"
	"log"
	"os"
)

func main() {
	aplicacao := app.Gerar()

	if err := aplicacao.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
