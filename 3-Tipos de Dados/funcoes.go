package main

import "fmt"

func main() {
	//var soma int
	fmt.Println("Essa é a funcao principal")
	soma := somar(8, 10)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println("funcao aninhada ", txt)
	}

	f("Olha texto aqui")

	//resultadoSoma, resultadoSubt := calculosMatematicos(10, 8)

	resultadoSoma, _ := calculosMatematicos(10, 8) // se eu não quiser usar um dos retornos coloca um _

	fmt.Println("Soma: ", resultadoSoma)

}

func somar(n1 int, n2 int) int { // Especifica o retorno depois do fechar parentese dos parametros
	fmt.Println("Essa é a funcao somar")
	return n1 + n2

}

func calculosMatematicos(n1, n2 int) (int, int) { // funcao com mais de um retorno, coloca entre parenteses todos os tipos de retorno
	soma := somar(n1, n2)
	subtracao := n1 - n2

	return soma, subtracao
}
