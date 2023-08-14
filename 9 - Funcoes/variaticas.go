package main

import "fmt"

func soma(numeros ...int) int {
	fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}
func main() {
	total := soma(1, 2, 3, 4, 5, 6, 103, 7, 9)
	fmt.Println(total)

	escrever("Oi", 1, 2, 3, 4, 5)
}
