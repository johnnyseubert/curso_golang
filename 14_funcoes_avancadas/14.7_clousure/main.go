package main

func clousure() func() {
	texto := "Dentro da clousure"

	funcao := func() {
		println(texto)
	}
	return funcao
}

func main() {
	novaFuncao := clousure()
	novaFuncao()
}
