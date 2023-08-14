package main

import "fmt"

func recuperaExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução Recuperada com Sucesso")
	}
}
func alunoAprovadro(n1, n2 float32) bool {
	defer recuperaExecucao()
	media := (n1 + n2) / 2
	defer fmt.Println("Media Calculada. Resultado será retornado")
	if media > 6 {
		fmt.Println("Aluno Aprovado")
		return true
	} else if media < 6 {
		fmt.Println("Aluno Reprovado")
		return false
	}

	panic("A média é exatamente 6")

}

func main() {

	fmt.Println(alunoAprovadro(6, 8))
	fmt.Println("Pós Execução")
}
