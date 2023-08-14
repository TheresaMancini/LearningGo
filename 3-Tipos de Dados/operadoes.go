package main

import (
	"fmt"
)

func main() {
	//Aritmeticos

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println("\nsoma: ", soma, "\nsubtracao: ", subtracao, "\ndvisao: ", divisao, "\nmultiplicacao", multiplicacao, "\nrestoDivisao: ", restoDivisao)

	// Relacionais
	fmt.Println("1 == 2", 1 == 2)
	fmt.Println("1 > 2", 1 > 2)
	fmt.Println("1 < 2", 1 < 2)

	//logicos
	verdadeiro, falso := true, false

	fmt.Println("&&", verdadeiro && verdadeiro)
	fmt.Println("&&", verdadeiro && falso)
	fmt.Println("||", verdadeiro || verdadeiro)
	fmt.Println("||", verdadeiro || falso)
	fmt.Println("!", !verdadeiro)
	fmt.Println("!", !falso)

	// unario
	numero := 10

	numero++
	fmt.Println("++", numero)
	numero += 2
	fmt.Println("+=2", numero)

}
