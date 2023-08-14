package main

import "fmt"

func main() {
	var variavel1 string = "Varável 1" // explicitando tipo da variavel
	variavel2 := "Variável 2"          // Se declarar e não usar, da erro de compilação - Declaracao implicita. Infere o tipo da variável pela sua inicialização

	var ( // Declarando mais de uma variavel de uma vez, explicitando os tipos
		variavel3 string
		variavel4 string
	)

	variavel5, variavel6 := "Variavel5", "Variavel 6"

	const constante string = "Textin"
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5, variavel6)
	fmt.Println(constante)

	// Para trocar valor de variaveis nao precisa de variavel auxiliar

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
