package physics

type Temperature struct {
	kelvin float64
}

func Celcius(temp float64) Temperature {
	return Temperature{temp + 273.15}
}

func Kelvin(temp float64) Temperature {
	return Temperature{temp}
}

func Farenehit(temp float64) Temperature {
	return Temperature{(5.0/9.0)*(temp-32) + 273.15}
}

func (t Temperature) Celcius() float64 {
	return t.kelvin - 273.15
}

func (t Temperature) Kelvin() float64 {
	return t.kelvin
}

func (t Temperature) Farenehit() float64 {
	return (t.kelvin-273.15)*1.8 + 32.0
}
