package main

import (
	"fmt"
)

func main() {

	i := 0
	fmt.Println(i)
	for i < 10 {
		i++
		fmt.Println("incrementando i ")
		//time.Sleep(time.Second)

	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("incrementando j ", j)
		//time.Sleep(time.Second)
	}

	nomes := [3]string{"Joao", "David", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	/// pode iterar sobre array, slice, string (retorna ascII), map. Iterando sempre traz primeiro o indice depois o valor. Se for um map e chave, valor
}
