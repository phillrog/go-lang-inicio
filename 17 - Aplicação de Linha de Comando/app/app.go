package App

import "github.com/urfave/cli"

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "busca IPs e Nomes de Servidor na Internet"

	return app
}
