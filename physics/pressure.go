package physics

type Pressure struct {
	pascals float64
}

func KiloPascals(pressure float64) Pressure {
	return Pressure{pressure * 1000}
}

func Pascals(pressure float64) Pressure {
	return Pressure{pressure}
}

func PoundsPerSquareInch(pressure float64) Pressure {
	return Pressure{pressure * 6894.757293}
}

func Bars(pressure float64) Pressure {
	return Pressure{pressure * 100000}
}

func Atmospheres(pressure float64) Pressure {
	return Pressure{pressure * 101325}
}

func (p Pressure) KiloPascals() float64 {
	return p.pascals / 1000
}

func (p Pressure) Pascals() float64 {
	return p.pascals
}

func (p Pressure) PoundsPerSquareInch() float64 {
	return p.pascals / 6894.757293
}

func (p Pressure) Bars() float64 {
	return p.pascals / 100000
}

func (p Pressure) Atmospheres() float64 {
	return p.pascals / 101325
}
