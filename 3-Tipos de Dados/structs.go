package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
	end   endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Structs")

	var u usuario

	u.nome = "Theresa"
	u.idade = 26

	enderecoExemplo := endereco{"endereco", 23}

	u2 := usuario{nome: "Hernan", idade: 29, end: enderecoExemplo}

	u3 := usuario{nome: "Marcia", idade: 58, end: endereco{"Rua", 23}}

	fmt.Println(u)
	fmt.Println(u2)
	fmt.Println(u3)
}
