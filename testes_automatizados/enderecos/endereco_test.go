package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(test *testing.T) {
	test.Parallel() // Diz que os testes que tem essa flag pode ser feito em paralelo caso queira

	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Rua Almirante Saldanha da Gama", retornoEsperado: "RUA"},
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "AVENIDA"},
		{enderecoInserido: "Estrada do Sol", retornoEsperado: "ESTRADA"},
		{enderecoInserido: "Rodovia BR470", retornoEsperado: "RODOVIA"},
		{enderecoInserido: "Praça das Rosas", retornoEsperado: "INVALID TYPE"},
		{enderecoInserido: "", retornoEsperado: "INVALID TYPE"},
	}

	for _, cenario := range cenariosDeTeste {
		resultado := TipoDeEndereco(cenario.enderecoInserido)

		if resultado != cenario.retornoEsperado {
			test.Errorf("O tipo recebido foi %s e o esperado é %s, para o valor %s",
				resultado,
				cenario.retornoEsperado,
				cenario.enderecoInserido,
			)
		}
	}
}

func TestTipoDeEndereco2(test *testing.T) {
	test.Parallel() // Diz que os testes que tem essa flag pode ser feito em paralelo caso queira

	cenariosDeTeste := []cenarioDeTeste{
		{enderecoInserido: "Rua Almirante Saldanha da Gama", retornoEsperado: "RUA"},
		{enderecoInserido: "Avenida Paulista", retornoEsperado: "AVENIDA"},
		{enderecoInserido: "Estrada do Sol", retornoEsperado: "ESTRADA"},
		{enderecoInserido: "Rodovia BR470", retornoEsperado: "RODOVIA"},
		{enderecoInserido: "Praça das Rosas", retornoEsperado: "INVALID TYPE"},
		{enderecoInserido: "", retornoEsperado: "INVALID TYPE"},
	}

	for _, cenario := range cenariosDeTeste {
		resultado := TipoDeEndereco(cenario.enderecoInserido)

		if resultado != cenario.retornoEsperado {
			test.Errorf("O tipo recebido foi %s e o esperado é %s, para o valor %s",
				resultado,
				cenario.retornoEsperado,
				cenario.enderecoInserido,
			)
		}
	}
}
