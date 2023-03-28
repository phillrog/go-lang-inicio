package app

import (
	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação
func Gerar() *cli.App {
	app := &cli.App{
		Name:  "Aplicação de Linha de Comando",
		Usage: "busca IPs e Nomes de Servidor na Internet",
	}

	return app
}
