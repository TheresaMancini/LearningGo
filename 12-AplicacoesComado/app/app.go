package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
) // para conseguir fazer o import, devo colocar no go.work na pasta mais externa. Pra poder usar os modulos em cada uma das pastas de aplicação

// Gerar vai retornar a aplicação de linha de compando pronta para ser executada
func Gerar() *cli.App { //cli é o nome do pacote. App é o tipo que ta dentro do pacote
	// funcao Gerar com G maiusculo pra ser uma funcao que pode ser acessada externamente
	app := cli.NewApp()
	app.Name = "App de Linha de Comando"                    // Nome da minha aplicacao
	app.Usage = "Busca Ips e Nomes de Servidor na Internet" // Uso da aplicacao

	flags := []cli.Flag{ // variavel flags que vai definir quais flags a aplicacao pode receber, ou o comando. Como as flags vao ser as mesmas para ambos os comandos, criando como variavel, se fosse diferente criar no proprio comando
		cli.StringFlag{
			Name:  "host", // definindo a flag --host
			Value: "devbook.com.br",
		},
	}

	app.Commands = []cli.Command{ // slice de comandos que a aplicacao vai entender cli.Commands é o tipo dos comandos. cli pq vem do pacote
		{
			Name:   "ip",
			Usage:  "Busca IPs de enderecos na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}
	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")
	// net
	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
