package forma

type Forma interface {
	//Obtém a área total da forma geométrica
	ObterArea() float64
	ObterPerimetro() float64
	String() string
}
