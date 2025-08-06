package forma

import "math"

type Retangulo struct {
	base, altura float64
}

func (c Retangulo) ObterArea() float64 {
	return c.base * c.altura
}

func (c Retangulo) ObterPerimetro() float64 {
	return math.Pow(c.base, 2.0) + math.Pow(c.altura, 2.0)
}

func (c Retangulo) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}
