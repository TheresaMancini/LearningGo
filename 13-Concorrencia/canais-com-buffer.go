package main

import "fmt"

func main() {
	canal := make(chan string, 2)

	canal <- "Olá Mundo" // Enviei o Dado fica travado aqui esperando alguem receber
	canal <- "Olá Mundo 2"
	//canal <- "Olá Mundo 3" // Se enviar um 3, excede o buffer e trava aqui
	mensagem := <-canal // Nunca chegou nessa linha
	mensagem2 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
