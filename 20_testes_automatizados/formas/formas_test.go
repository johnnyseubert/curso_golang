package formas

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		retangulo := Retangulo{Altura: 10, Largura: 12}

		areaEsperada := float64(120)
		areaCalculada := retangulo.Area()

		if areaCalculada != areaEsperada {
			t.Fatalf("Area calculada incorreta. Esperado: %.2f, Calculado: %.2f", areaEsperada, areaCalculada)
		}

	})

	t.Run("Circulo", func(t *testing.T) {
		circulo := Circulo{Raio: 10}

		areaEsperada := float64(math.Pi * 100)
		areaCalculada := circulo.Area()

		if areaCalculada != areaEsperada {
			t.Fatalf("Area calculada incorreta. Esperado: %.2f, Calculado: %.2f", areaEsperada, areaCalculada)
		}

	})

}
