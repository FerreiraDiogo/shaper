package forma

import "math"

// Retangulo representa um retângulo com base e altura.
type Retangulo struct {
	Base, Altura float64
}

// ObterArea retorna a área do retângulo.
func (r Retangulo) ObterArea() float64 {
	return r.Base * r.Altura
}

// ObterPerimetro retorna a soma dos quadrados da base e altura do retângulo.
func (r Retangulo) ObterPerimetro() float64 {
	return math.Pow(r.Base, 2.0) + math.Pow(r.Altura, 2.0)
}

// String retorna uma representação textual do retângulo.
func (r Retangulo) String() string {
	return "I just dont know what to do with myself...[rectangle]"
}
