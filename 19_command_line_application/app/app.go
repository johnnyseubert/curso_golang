package app

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// Vai retornar a aplicação CLI pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca ips e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca um ip na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome do servidor na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		panic(err)
	}

	for _, ip := range ips {
		fmt.Println("HOST", host, "IP", ip.String())
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, err := net.LookupNS(host)

	if err != nil {
		panic(err)
	}

	for _, servidor := range servidores {
		fmt.Println("HOST", host, "SERVIDOR", servidor.Host)
	}

}
