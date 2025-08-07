package forma

import (
	"math"
	"testing"
)

func TestRetanguloObterArea(t *testing.T) {
	r := Retangulo{base: 4, altura: 5}
	expected := 20.0
	area := r.ObterArea()
	if area != expected {
		t.Errorf("ObterArea() = %v; want %v", area, expected)
	}
}

func BenchmarkRetanguloObterArea(b *testing.B) {
	r := Retangulo{base: 1.0, altura: 1.0}
	for i := 0; i < b.N; i++ {
		r.ObterArea()
	}
}

func TestRetanguloObterPerimetro(t *testing.T) {
	r := Retangulo{base: 3, altura: 4}
	expected := math.Pow(3, 2.0) + math.Pow(4, 2.0)
	perimetro := r.ObterPerimetro()
	if perimetro != expected {
		t.Errorf("ObterPerimetro() = %v; want %v", perimetro, expected)
	}
}

func BenchmarkRetanguloObterPerimetro(b *testing.B) {
	r := Retangulo{base: 1.0, altura: 1.0}
	for i := 0; i < b.N; i++ {
		r.ObterPerimetro()
	}
}

func TestRetanguloString(t *testing.T) {
	r := Retangulo{}
	expected := "I just dont know what to do with myself...[rectangle]"
	if r.String() != expected {
		t.Errorf("String() = %q; want %q", r.String(), expected)
	}
}

func BenchmarkRetanguloString(b *testing.B) {
	r := Retangulo{base: 1.0, altura: 1.0}
	for i := 0; i < b.N; i++ {
		r.String()
	}
}
