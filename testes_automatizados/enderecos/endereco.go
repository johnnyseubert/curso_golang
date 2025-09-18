package enderecos

import (
	"slices"
	"strings"
)

func TipoDeEndereco(enderecoCompleto string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	primeiraPalavraDoEndereco := strings.Split(enderecoCompleto, " ")[0]
	enderecoEmMinuscula := strings.ToLower(primeiraPalavraDoEndereco)

	enderecoTemUmTipoValido := slices.Contains(tiposValidos, enderecoEmMinuscula)

	if enderecoTemUmTipoValido {
		return strings.ToTitle(enderecoEmMinuscula)
	}

	return "INVALID TYPE"
}
