package forma

type Triangulo struct {
	base, altura, face float64
}

func (t Triangulo) ObterArea() float64 {
	return (t.base * t.altura) / 2
}

func (t Triangulo) ObterPerimetro() float64 {
	return t.face * 3
}

func (t Triangulo) String() string {
	return "I just dont know what to do with myself.[triangulo]]"
}
