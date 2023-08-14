package main

import "fmt"

func main() {
	tarefas := make(chan int, 45) // Buffer de 45, pode receber 45 inteiros
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for j := 0; j < 45; j++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	// Padrao woorker pool
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibonacci(n-2) + fibonacci(n-1)
	}
}
