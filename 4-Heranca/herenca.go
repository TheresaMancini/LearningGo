package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{nome: "Theresa", sobrenome: "Mancini", idade: 26}

	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "UFRN"}

	fmt.Println(e1)

	e2 := estudante{pessoa: p1, curso: "Engenharia", faculdade: "UFRN"}

	fmt.Println(e2)

	fmt.Println(e2.nome)

	e3 := estudante{pessoa: pessoa{nome: "Maria"}, curso: "Engenharia", faculdade: "UFRN"}

	fmt.Println(e3.nome)
}
