package forma

type Forma interface {
	//Obtém a área total da forma geométrica
	ObterArea() float64
	ObterPerimetro() float64
	String() string
}

type ByArea []Forma

func (a ByArea) Len() int {
	return len(a)
}

func (a ByArea) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByArea) Less(i, j int) bool {
	return a[i].ObterArea() < a[j].ObterArea()
}
