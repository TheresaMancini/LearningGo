package main

import (
	"fmt"
	"time"
)

func main() {
	// Usando CHAIN CANAIS

	canal := make(chan string)

	go escreverCanal("Olá Mundo", canal)

	for {
		mensagem, aberto := <-canal // Canal ta esperando receber o valor.

		if !aberto { //evitando o deadlock
			break
		}
		fmt.Println(mensagem)
	}

	// USANDO WAIT GROUP
	/*var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá Mundo")
		waitGroup.Done() // 2-1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // 1-1
	}()

	waitGroup.Wait() // 0

	fmt.Println("Só depois que terminar o waitgroup")*/

	// USANDO GO ROTINE
	/*go escrever("Olá Mundo") //goroutine
	escrever("Programando em Go")*/
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func escreverCanal(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		//fmt.Println(texto)
		canal <- texto /// passando o valor para o canal. Canal recebe esse texto
		time.Sleep(time.Second)
	}
	close(canal) // fechando o canal, assim não envia nem recebe mais dados
}
