package forma

import "math"

// Poligono representa um polígono regular com número de faces e tamanho da face.
type Poligono struct {
	NumeroDeFaces int
	TamanhoDaFace float64
}

// ObterArea retorna a área do polígono regular.
func (p Poligono) ObterArea() float64 {
	return (p.ObterPerimetro() / 2) * p.calcularApotema()
}

// ObterPerimetro retorna o perímetro do polígono regular.
func (p Poligono) ObterPerimetro() float64 {
	return p.TamanhoDaFace * float64(p.NumeroDeFaces)
}

// String retorna uma representação textual do polígono.
func (p Poligono) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}

func (p Poligono) calcularApotema() float64 {
	return (p.TamanhoDaFace / 2) * math.Atanh(math.Pi/float64(p.NumeroDeFaces))
}
