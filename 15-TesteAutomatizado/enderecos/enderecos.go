package enderecos

import (
	"strings"
)

// TipoDeEndereco verifia se um endereco tem um tio válido e o retorna
func TipoDeEndereco(endereco string) string {
	// Minha funcao aplicao.
	// Funcao que sera testada.
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"} // slice

	//convertendo endereco todo pra letra minuscula

	enderecoMinuscula := strings.Title(endereco)

	tipoValido := false

	// pegando primeira palavra do endereco

	primeiraPalavra := strings.Split(enderecoMinuscula, " ")[0] // Pegando o primeiro elemento do slice criado a partir da string endereco, que foi separado usando o " " como separador
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			tipoValido = true
		}
	}

	if tipoValido {
		return primeiraPalavra
	}

	return "Tipo Invalido"
}
