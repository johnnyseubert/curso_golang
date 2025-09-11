package main

func main() {
	numero := 2

	if numero > 2 {
		println("Maior que 2")
	} else if numero == 2 {
		println("Igual a 2")
	} else {
		println("Menor que 2")
	}

	// Cria a variavel dentro do if else etc.. e ja faz a verificação se for verdadeira entra no if
	if outroNumero := 3; outroNumero == 2 {
		println(outroNumero)
	} else {
		println(outroNumero)
	}

	// Fora do if nao existe a variavel outroNumero
	// println(outroNumero) // ERRO
}
