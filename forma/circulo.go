package forma

import "math"

type Circulo struct {
	Raio float64
}

func (c Circulo) ObterArea() float64 {
	return math.Pi * math.Pow(c.Raio, 2.0)
}

func (c Circulo) ObterPerimetro() float64 {
	return 2.0 * math.Pi * c.Raio
}

func (c Circulo) String() string {
	return "I just dont know what to do with myself..."
}
