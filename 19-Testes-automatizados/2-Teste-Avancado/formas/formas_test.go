package main

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Retângulo", func(t *testing.T) {
		ret := Retangulo{10, 12}
		areaEsperado := float64(120)
		areaRecebida := ret.Area()

		if areaEsperado != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f ", areaRecebida, areaEsperado)
		}
	})

	t.Run("Círculo", func(t *testing.T) {
		circulo := Circulo{10}

		areaEsperado := float64(math.Pi * 100)
		areaRecebida := circulo.Area()

		if areaEsperado != areaRecebida {
			t.Fatalf("A área recebida %f é diferente da área esperada %f ", areaRecebida, areaEsperado)
		}
	})
}
