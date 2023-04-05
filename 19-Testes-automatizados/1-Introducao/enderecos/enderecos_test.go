package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenarioDeTeste := []cenarioDeTeste{
		{"Rua Pestallozi", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia bandeirantes", "Rodovia"},
		{"", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTeste {
		retornoRecebido := TipoEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido  %sé diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
