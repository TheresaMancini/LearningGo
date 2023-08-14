package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Olá Mundo")
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
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
