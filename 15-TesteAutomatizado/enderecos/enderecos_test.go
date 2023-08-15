// Teste Unitario
// Testa a menor unidade do código. Testa Apenas uma funcao
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado. Recebeu %s e esperava %s", tipoDeEnderecoRecebido, tipoDeEnderecoEsperado)
	} // se cair no erro é diferente do esperado, o teste deu erro
}
