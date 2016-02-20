package physics

type Mass struct {
	grams float64
}

func Kilograms(mass float64) Mass {
	return Mass{mass * 1000}
}

func Grams(mass float64) Mass {
	return Mass{mass}
}

func Pounds(mass float64) Mass {
	return Mass{mass * 453.59237}
}

func Ounces(mass float64) Mass {
	return Mass{mass * 28.349523125}
}

func (m Mass) Kilograms() float64 {
	return m.grams / 1000.0
}

func (m Mass) Grams() float64 {
	return m.grams
}

func (m Mass) Pounds() float64 {
	return m.grams / 453.59237
}

func (m Mass) Ounces() float64 {
	return m.grams / 28.349523125
}
