package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoTeste := "Rua Pestallozzi"
	tipoEnderecoEsperado := "Rua"
	tipoEnderecoValidar := TipoEndereco(enderecoTeste)

	if tipoEnderecoValidar != tipoEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s mas recebeu %s", tipoEnderecoEsperado, tipoEnderecoValidar)
	}
}
