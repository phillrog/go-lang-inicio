package enderecos

import "strings"

// TipoEndereco verifica se o endereço é válido
func TipoEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}

	primeiraPavraDoEndereco := strings.Split(endereco, " ")[0]

	enderecoValido := false

	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == strings.ToLower(primeiraPavraDoEndereco) {
			enderecoValido = true
		}
	}

	if enderecoValido {
		return primeiraPavraDoEndereco
	}

	return "Tipo Inválido"
}
