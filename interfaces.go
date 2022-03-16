package figuras

import "fmt"

type Geometrica interface {
	area() float64
	perimetro() float64
}

//nombre de metodo con mayusculas es PUBLICO
func Medidas(g Geometrica) {
	fmt.Println("medidas: ", g)
	fmt.Println("area: ", g.area())
	fmt.Println("perimetro: ", g.perimetro())
}
