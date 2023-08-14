package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	// ponteiros

	var variavel1 int = 10

	var variavel2 int = variavel1

	fmt.Println("v1: ", variavel1, " v2: ", variavel2)

	variavel1++

	fmt.Println("v1 ++")

	fmt.Println("v1: ", variavel1, " v2: ", variavel2)
	// Atribuindo valor usando ponteiro

	fmt.Println("Atribuindo valor usando ponteiro")

	var variavel3 int
	var ponteiro *int

	variavel3 = 100

	fmt.Println("v3: ", variavel3, " ponteiro: ", ponteiro)

	ponteiro = &variavel3

	fmt.Println("ponteiro = &variavel3")

	fmt.Println("v3: ", variavel3, " ponteiro: ", ponteiro)

	fmt.Println("v3: ", variavel3, " *ponteiro: ", *ponteiro) /// desreferenciacao

	variavel3++

	fmt.Println("variavel3++")
	fmt.Println("v3: ", variavel3, " *ponteiro: ", *ponteiro) /// desreferenciacao

}
