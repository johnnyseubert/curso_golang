package main

import (
	"errors"
	"fmt"
)

func main() {
	var textoDefault string
	var nome string = "Johnny"
	fmt.Println("STRINGS: ", textoDefault, nome)

	var boleanoDefault bool
	var possui bool = true
	fmt.Println("BOOLEAN: ", boleanoDefault, possui)

	var inteiroDefault int
	var inteiro int = 0                       // platform dependent: 32 or 64 bits
	var inteiro8 int8 = 127                   // -128 to 127
	var inteiro16 int16 = 32767               // -32768 to 32767
	var inteiro32 int32 = 2147483647          // -2147483648 to 2147483647
	var inteiro64 int64 = 9223372036854775807 // -9223372036854775808 to 9223372036854775807
	fmt.Println("INTEIROS: ", inteiroDefault, inteiro, inteiro8, inteiro16, inteiro32, inteiro64)

	// Somente numeros positivos
	var inteiroPositivoDefault uint                     // platform dependent: 32 or 64 bits
	var inteiroPositivo8 uint8 = 255                    // 0 to 255
	var inteiroPositivo16 uint16 = 65535                // 0 to 65535
	var inteiroPositivo32 uint32 = 4294967295           // 0 to 4294967295
	var inteiroPositivo64 uint64 = 18446744073709551615 // 0 to 18446744073709551615
	fmt.Println("INTEIROS POSITIVOS: ", inteiroPositivoDefault, inteiroPositivo8, inteiroPositivo16, inteiroPositivo32, inteiroPositivo64)

	var flutuanteDefault float64
	var flutuante float64 = 0.0
	fmt.Println("FLUTUANTES: ", flutuanteDefault, flutuante)

	// Rune é um tipo para char com apenas um digito e ele nao retorna o "a" ele retorna o valor inteiro do "a" na tabela asc
	var runaDefault rune
	var runa rune = 'a'
	fmt.Println("RUNAS: ", runaDefault, runa)

	var biteDefault byte // alias para uint8
	var bite byte = 255
	fmt.Println("BYTES: ", biteDefault, bite)

	var erroDefault error
	var erro error = errors.New("Erro interno")
	fmt.Println("ERROR: ", erroDefault, erro)

}
