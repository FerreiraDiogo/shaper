package forma

import "math"

type Poligono struct {
	numeroDeFaces int
	tamanhoDaFace float64
}

func (p Poligono) ObterArea() float64 {
	return (p.ObterPerimetro() / 2) * p.calcularApotema()
}

func (p Poligono) ObterPerimetro() float64 {
	return p.tamanhoDaFace * float64(p.numeroDeFaces)
}

func (p Poligono) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}

func (p Poligono) calcularApotema() float64 {
	return (p.tamanhoDaFace / 2) * math.Atanh(math.Pi/float64(p.numeroDeFaces))
}
