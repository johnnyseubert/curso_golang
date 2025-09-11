package main

func main() {
	// No go só existe o for, nao existe while do while forEach etc...

	// Tradicional
	for i := 0; i > -200; i-- {
		println("Tradicional", i)
	}

	// Equivalente a um forEach
	for i := range 5 {
		println("Range Default", i)
	}

	// Equivalente a um forEach com array
	meuRange := []int{10, 20, 30, 40, 50}
	for i, valor := range meuRange {
		println("Meu Range", i, valor)
	}

	// Equivalente a um while
	indice := 0
	for indice < 10 {
		indice++
		// time.Sleep(time.Millisecond * 500)
		println("While", indice)
	}

	// Equivalente a um while(true)
	for {
		println("Infinito")
		break // Só vai parar o for se tiver um break equivalente a um while(true) em outras linguagens
	}

	for indice, letra := range "PALAVRA" {
		println("String", indice, letra, string(letra))
	}

	mapper := map[string]string{
		"nome":      "Gustavo",
		"sobrenome": "Guanabara",
	}

	for chave, valor := range mapper {
		println("CHAVE", chave, "VALOR", valor)
	}

}
