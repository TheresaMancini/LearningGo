package main

import (
	"errors"
	"fmt"
)

func main() {
	// Tipos de dados basicos no go
	//Inteiros
	//int8, int16, int32, int64 int32 tem o alias de rune
	//usa o rune quando são variaveis que estao se referiando a caracteres alfanumericos
	//uint8 é o byte.
	// int o int sozinha assume a arquitetura do computador como base.

	//uint - unsigned int. Int sem sinal
	var numero1 uint = 132
	var numero2 int32 = 1000000
	fmt.Println(numero1)
	fmt.Println(numero2)

	// Reais

	//float32 float64
	var real1 float32 = 123.54
	var real2 float64 = 1555.89

	fmt.Println(real1)
	fmt.Println(real2)

	//String cadeia de caracteres
	var str string = "Texto"
	fmt.Println(str)

	//No go nao tem char. O que seria o char vai ser transformado em um numero.
	char := 'B'
	fmt.Println(char) // imprime o inteiro que corresponde ao valor na tabela ASC desse caracter

	// Booleano - true e false
	var booleano1 bool = true
	var booleano2 bool
	fmt.Println(booleano1)
	fmt.Println(booleano2)
	//erro - O tratamento de erro no Go é diferente do convencional

	var erro error
	fmt.Println(erro)

	//para criar um erro tem que usar o pacote errors
	var errocriado error = errors.New("Erro Interno")
	fmt.Println(errocriado)
}
