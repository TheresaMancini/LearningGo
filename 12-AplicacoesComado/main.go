package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de Partida")
	//o main só vai executar a aplicação

	aplicacao := app.Gerar() //app eh o nome do pacote que acabou de ser criado

	erro := aplicacao.Run(os.Args) // os.Args parametro para pegar os comandos do sistema operacional da linha de comando

	if erro != nil {
		log.Fatal(erro) // Fatal para a aplciasao
	}
}
