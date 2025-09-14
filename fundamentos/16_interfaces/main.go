package main

import (
	"fmt"
	"math"
)

// A aplicação da interface é feita de forma implícita, ou seja, não é necessário declarar que uma struct implementa uma interface
// Ao criar o metodo que pertence a interface, a struct já implementa a interface
// Uma struct pode implementar várias interfaces
// Uma interface pode ser implementada por várias structs
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// Ao comentar essa funcao o Circulo nao implementa mais a interface forma ai da erro
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma e: %.2f\n", f.area())
}

func main() {
	r := retangulo{altura: 10, largura: 15}
	c := circulo{raio: 10}

	escreverArea(r)
	escreverArea(c)

}
