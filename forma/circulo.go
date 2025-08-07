package forma

import "math"

// Circulo representa um círculo com determinado raio.
type Circulo struct {
	Raio float64
}

// ObterArea retorna a área do círculo.
func (c Circulo) ObterArea() float64 {
	return math.Pi * math.Pow(c.Raio, 2.0)
}

// ObterPerimetro retorna o perímetro do círculo.
func (c Circulo) ObterPerimetro() float64 {
	return 2.0 * math.Pi * c.Raio
}

// String retorna uma representação textual do círculo.
func (c Circulo) String() string {
	return "I just dont know what to do with myself..."
}
