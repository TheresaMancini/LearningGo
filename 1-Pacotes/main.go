package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da funcao main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("theresa.mancini@gmail.com")
	fmt.Println(erro)
}
