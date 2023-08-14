package main

import "fmt"

func funcao1() {
	fmt.Println("Executando funcao 1")
}

func funcao2() {
	fmt.Println("Executando funcao 2")
}

func alunoAprovadro(n1, n2 float32) bool {

	media := (n1 + n2)
	defer fmt.Println("Media Calculada. Resultado será retornado")
	if media >= 6 {
		fmt.Println("Aluno Aprovado")
		return true
	} else {
		fmt.Println("Aluno Reprovado")
		return false
	}

}

func main() {
	defer funcao1()
	funcao2()

	fmt.Println(alunoAprovadro(5, 6))
}
