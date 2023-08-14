package main

import "fmt"

func main() {
	func() {
		fmt.Println("Olá sou uma funcao anonima")
	}()
	// para executar a funcao que é anonima faz ()

	func(texto string) {
		fmt.Println("Olá sou uma funcao anonima com parametro ", texto)
	}("Uau")
	// para executar a funcao que é anonima faz ()
}
