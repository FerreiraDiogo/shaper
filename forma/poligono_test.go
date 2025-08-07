package forma

import (
	"testing"
)

func TestPoligonoObterArea(t *testing.T) {
	p := Poligono{numeroDeFaces: 6, tamanhoDaFace: 2}
	expected := (p.ObterPerimetro() / 2) * p.calcularApotema()
	area := p.ObterArea()
	if area != expected {
		t.Errorf("ObterArea() = %v; want %v", area, expected)
	}
}

func TestPoligonoObterPerimetro(t *testing.T) {
	p := Poligono{numeroDeFaces: 5, tamanhoDaFace: 3}
	expected := 5 * 3.0
	perimetro := p.ObterPerimetro()
	if perimetro != expected {
		t.Errorf("ObterPerimetro() = %v; want %v", perimetro, expected)
	}
}

func TestPoligonoString(t *testing.T) {
	p := Poligono{}
	expected := "I just dont know what to do with myself...[rectangle]"
	if p.String() != expected {
		t.Errorf("String() = %q; want %q", p.String(), expected)
	}
}

func BenchmarkPoligonoObterArea(b *testing.B) {
	p := Poligono{numeroDeFaces: 6, tamanhoDaFace: 2}
	for i := 0; i < b.N; i++ {
		p.ObterArea()
	}
}

func BenchmarkPoligonoObterPerimetro(b *testing.B) {
	p := Poligono{numeroDeFaces: 5, tamanhoDaFace: 3}
	for i := 0; i < b.N; i++ {
		p.ObterPerimetro()
	}
}

func BenchmarkPoligonoString(b *testing.B) {
	p := Poligono{numeroDeFaces: 3, tamanhoDaFace: 1}
	for i := 0; i < b.N; i++ {
		p.String()
	}
}
