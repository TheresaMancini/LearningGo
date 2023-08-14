package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Praça Senador Mariz")
	fmt.Println(tipoEndereco)
}
