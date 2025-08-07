package forma

import "math"

type Retangulo struct {
	base, altura float64
}

func (r Retangulo) ObterArea() float64 {
	return r.base * r.altura
}

func (r Retangulo) ObterPerimetro() float64 {
	return math.Pow(r.base, 2.0) + math.Pow(r.altura, 2.0)
}

func (r Retangulo) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}
