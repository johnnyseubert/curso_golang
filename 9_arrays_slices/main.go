package main

import "fmt"

func main() {
	// Array com tamanho fixo
	var array1 [5]string
	array1[0] = "Posição 0"
	fmt.Println(array1)

	// Array com inicialização direta
	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)

	// Array com tamanho dinâmico baseado na quantidade de elementos inicializados
	array3 := [...]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5", "Posição 6"}
	fmt.Println(array3)

	// Slice (tamanho dinâmico)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("SLICE", slice)

	slice2 := array2[2:3] // Slice a partir do array2, do índice 2 ao 3 (3 não incluso)
	fmt.Println("SLICE2", slice2)

	slice3 := array2[:] // Slice a partir do array2, do início ao fim
	fmt.Println("SLICE3", slice3)

	slice4 := array2[3:] // Slice a partir do array2, do índice 3 até o fim
	fmt.Println("SLICE4", slice4)

	// Adicionando um elemento no slice
	slice = append(slice, 6)
	fmt.Println("SLICE", slice)

	// Atualizando um valor no slice
	slice[0] = 100
	fmt.Println("SLICE", slice)

	// Removendo um elemento no slice (removendo o elemento do índice 2)
	slice = append(slice[:2], slice[3:]...)
	fmt.Println("SLICE", slice)

	fmt.Println("")

	// ARRAYS INTERNOS
	slice5 := make([]int, 1, 3) // Cria um slice de int com tamanho 10 e capacidade 11

	fmt.Println("Tamanho antes de adicionar", len(slice5))    // Tamanho do slice
	fmt.Println("Capacidade antes de adicionar", cap(slice5)) // Capacidade do slice

	slice5 = append(slice5, 1)
	slice5 = append(slice5, 2)
	slice5 = append(slice5, 3)
	// Quando atinge a capacidade máxima, o Go cria um novo array interno com o dobro da capacidade

	fmt.Println(slice5)
	fmt.Println("Tamanho", len(slice5))    // Tamanho do slice
	fmt.Println("Capacidade", cap(slice5)) // Capacidade do slice
}
