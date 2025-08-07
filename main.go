package main

import (
	"fmt"
	"main/forma"
	"sort"
)

var formas []forma.Forma
var circulo,
	retangulo,
	triangulo,
	poligono forma.Forma

func init() {
	circulo = forma.Circulo{Raio: 4.0}
	retangulo = forma.Retangulo{Base: 2.0, Altura: 2.0}
	triangulo = forma.Triangulo{Base: 2.0, Altura: 2.0, Face: 2.0}
	poligono = forma.Poligono{NumeroDeFaces: 6, TamanhoDaFace: 6.0}

	formas = make([]forma.Forma, 0)
	formas = append(formas, circulo)
	formas = append(formas, retangulo)
	formas = append(formas, triangulo)
	formas = append(formas, poligono)

}

func main() {
	fmt.Printf("A área total das formas no slice é %f\n", obterAreaTotal())
	fmt.Printf("a forma com a maior area é %s", obterMaiorArea())
}

// Obtem a Forma com a maior area
func obterMaiorArea() forma.Forma {
	sort.Sort(forma.ByArea(formas))
	return formas[len(formas)-1]
}

// obterAreaTotal calcula e retorna a área de todas as formas no slice
func obterAreaTotal() float64 {
	var areaTotal float64
	for _, v := range formas {
		areaTotal += v.ObterArea()
	}
	return areaTotal
}
