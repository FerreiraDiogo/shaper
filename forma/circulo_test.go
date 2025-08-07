package forma

import (
	"math"
	"strings"
	"testing"
)

func TestObterArea(t *testing.T) {
	//struct declarando os cenários de teste
	tests := []struct {
		raio float64
		want float64
	}{
		{1.0, math.Pi * math.Pow(1.0, 2.0)},
		{0.0, 0.0},
		{1.5, math.Pi * math.Pow(1.5, 2.0)},
		{-1.0, math.Pi * math.Pow(-1.0, 2.0)},
	}

	//iterando por cada cenário de teste
	for _, tt := range tests {
		c := Circulo{Raio: tt.raio}
		got := c.ObterArea()
		if got != tt.want {
			t.Errorf("Erro em Circulo.ObterPerimetro: got: %f, want %f", got, tt.want)
		}
	}
}
func TestObterPerimetro(t *testing.T) {
	tests := []struct {
		raio float64
		want float64
	}{
		{1.0, 2.0 * math.Pi * 1.0},
		{0.0, 0.0},
		{2.5, 2.0 * math.Pi * 2.5},
		{-3.0, 2.0 * math.Pi * -3.0},
	}

	for _, tt := range tests {
		c := Circulo{Raio: tt.raio}
		got := c.ObterPerimetro()
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("ObterPerimetro() with raio=%f: got %f, want %f", tt.raio, got, tt.want)
		}
	}
}

func TestString(t *testing.T) {
	c := Circulo{Raio: 1.0}
	mensagem := strings.TrimSpace(c.String())

	if mensagem == "" {
		t.Error("[Circulo] mensagem vazia")
	}
}

func TestCirculoObterArea(t *testing.T) {
	c := Circulo{Raio: 3}
	expected := math.Pi * math.Pow(3, 2.0)
	area := c.ObterArea()
	if area != expected {
		t.Errorf("ObterArea() = %v; want %v", area, expected)
	}
}

func TestCirculoObterPerimetro(t *testing.T) {
	c := Circulo{Raio: 3}
	expected := 2.0 * math.Pi * 3
	perimetro := c.ObterPerimetro()
	if perimetro != expected {
		t.Errorf("ObterPerimetro() = %v; want %v", perimetro, expected)
	}
}

func TestCirculoString(t *testing.T) {
	c := Circulo{}
	expected := "I just dont know what to do with myself..."
	if c.String() != expected {
		t.Errorf("String() = %q; want %q", c.String(), expected)
	}
}

func BenchmarkCirculoObterArea(b *testing.B) {
	c := Circulo{Raio: 1.0}
	for i := 0; i < b.N; i++ {
		c.ObterArea()
	}
}

func BenchmarkCirculoObterPerimetro(b *testing.B) {
	c := Circulo{Raio: 1.0}
	for i := 0; i < b.N; i++ {
		c.ObterPerimetro()
	}
}

func BenchmarkCirculoString(b *testing.B) {
	c := Circulo{Raio: 1.0}
	for i := 0; i < b.N; i++ {
		c.String()
	}
}
