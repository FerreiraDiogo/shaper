package forma

// Triangulo representa um triângulo com base, altura e face.
type Triangulo struct {
	Base, Altura, Face float64
}

// ObterArea retorna a área do triângulo.
func (t Triangulo) ObterArea() float64 {
	return (t.Base * t.Altura) / 2
}

// ObterPerimetro retorna o perímetro do triângulo equilátero.
func (t Triangulo) ObterPerimetro() float64 {
	return t.Face * 3
}

// String retorna uma representação textual do triângulo.
func (t Triangulo) String() string {
	return "I just dont know what to do with myself.[triangulo]]"
}
