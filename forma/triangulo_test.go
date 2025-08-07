package forma

import (
	"testing"
)

func TestTrianguloObterArea(t *testing.T) {
	tri := Triangulo{base: 10, altura: 5}
	expected := 25.0
	area := tri.ObterArea()
	if area != expected {
		t.Errorf("ObterArea() = %v; want %v", area, expected)
	}
}

func TestTrianguloObterPerimetro(t *testing.T) {
	tri := Triangulo{face: 7}
	expected := 21.0
	perimetro := tri.ObterPerimetro()
	if perimetro != expected {
		t.Errorf("ObterPerimetro() = %v; want %v", perimetro, expected)
	}
}

func TestTrianguloString(t *testing.T) {
	tri := Triangulo{}
	expected := "I just dont know what to do with myself.[triangulo]]"
	if tri.String() != expected {
		t.Errorf("String() = %q; want %q", tri.String(), expected)
	}
}

func BenchmarkTrianguloObterArea(b *testing.B) {
	tri := Triangulo{base: 1.0, altura: 1.0}
	for i := 0; i < b.N; i++ {
		tri.ObterArea()
	}
}

func BenchmarkTrianguloObterPerimetro(b *testing.B) {
	tri := Triangulo{face: 1.0}
	for i := 0; i < b.N; i++ {
		tri.ObterPerimetro()
	}
}

func BenchmarkTrianguloString(b *testing.B) {
	tri := Triangulo{base: 1.0, altura: 1.0, face: 1.0}
	for i := 0; i < b.N; i++ {
		tri.String()
	}
}
