package enderecos_test

import (
	. "go-lang-inicio/19-Testes-automatizados/1-Introducao/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
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

func TestSOma(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste falhou!")
	}
}
