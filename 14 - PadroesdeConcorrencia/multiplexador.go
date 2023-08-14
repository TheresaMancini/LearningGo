package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexar(escrever("Olá"), escrever("Programar"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)
	// recebe dois canais e retorna 1 canal

	go func() {
		for {
			select { // verificar qua canal tem mensagem
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida

}

func escrever(texto string) <-chan string { // recebe texto retorna um canal de string

	canal := make(chan string) // criando canal de string

	go func() { // go routine -- funcao anonima
		for {
			canal <- fmt.Sprintf("Valor recebido %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
