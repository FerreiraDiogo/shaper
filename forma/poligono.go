package forma

import "math"

type Poligono struct {
	numeroDeFaces int
	tamanhoDaFace float64
}

func (c Poligono) ObterArea() float64 {
	return (c.ObterPerimetro() / 2) * c.calcularApotema()
}

func (c Poligono) ObterPerimetro() float64 {
	return c.tamanhoDaFace * float64(c.numeroDeFaces)
}

func (c Poligono) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}

func (c Poligono) calcularApotema() float64 {
	return (c.tamanhoDaFace / 2) * math.Atanh(math.Pi/float64(c.numeroDeFaces))
}
