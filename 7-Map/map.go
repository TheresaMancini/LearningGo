package main

import "fmt"

func main() {
	fmt.Println("Map")

	usuario := map[string]string{"nome": "Pedro", "sobrenome": "Silva"}
	// chave nome, sobrenome
	// valores Pedro, Silva
	fmt.Println(usuario)

	fmt.Println(usuario["nome"]) //acessando uma chave específica

	usuario2 := map[string]map[string]uint{
		"map1": {
			"map2valor1": 2,
			"map2valor2": 3,
		},
		"map2": {
			"map2valor1": 2,
		},
	}

	fmt.Println(usuario2)

	// para deletar uma chave

	delete(usuario2, "map1")

	fmt.Println(usuario2)
}
