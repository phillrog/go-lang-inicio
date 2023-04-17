package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação
func Gerar() *cli.App {
	app := &cli.App{
		Name:  "Aplicação de Linha de Comando",
		Usage: "busca IPs e Nomes de Servidor na Internet",
		Commands: []*cli.Command{
			{
				Name:  "ip",
				Usage: "busca IPs de endereços na internet",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "devbook.com.br",
					},
				},
				Action: buscarIps,
			},
			{
				Name:  "servidores",
				Usage: "busca o nome dos servidores na internet",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "host",
						Value: "devbook.com.br",
					},
				},
				Action: buscarServidores,
			},
		},
	}

	return app
}

func buscarIps(ctx *cli.Context) error {
	host := ctx.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

	return erro
}

func buscarServidores(ctx *cli.Context) error {
	host := ctx.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

	return erro
}
