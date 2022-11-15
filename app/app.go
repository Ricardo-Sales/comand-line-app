package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IP e nome de servidor na internet"
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na Internet",
			Flags:  flags,
			Action: BuscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: BuscarServidores,
		},
	}

	return app

}

func BuscarIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
func BuscarServidores(c *cli.Context) {
	host := c.String("host")

	servs, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, serv := range servs {
		fmt.Println(serv.Host)
	}
}
