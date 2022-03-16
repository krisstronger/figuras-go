package figuras

import "math"

type Circulo struct {
	Radio float64
}

//metodo privado en minuscula
func (cir *Circulo) area() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

func (cir *Circulo) perimetro() float64 {
	return 2 * math.Pi * cir.Radio
}
