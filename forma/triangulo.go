package forma

type Triangulo struct {
	base, altura, face float64
}

func (c Triangulo) ObterArea() float64 {
	return (c.base * c.altura) / 2
}

func (c Triangulo) ObterPerimetro() float64 {
	return c.face * 3
}

func (c Triangulo) String() string {
	return "I just dont know what to do with myself.[triangulo]]"
}
