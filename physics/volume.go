package physics

type Volume struct {
	litres float64
}

func Milliliters(volume float64) Volume {
	return Volume{volume / 1000}
}

func Litres(volume float64) Volume {
	return Volume{volume}
}

func USGallons(volume float64) Volume {
	return Volume{volume * 3.785411784}
}

func USLiquidOunces(volume float64) Volume {
	return Volume{volume * 0.0295735295625}
}

func CubicMeters(volume float64) Volume {
	return Volume{volume * 1000}
}

func CubicInches(volume float64) Volume {
	return Volume{volume * 0.016387064}
}

func CubicFeet(volume float64) Volume {
	return Volume{volume * 28.316846592}
}

func (v Volume) Milliliters() float64 {
	return v.litres * 1000
}

func (v Volume) Litres() float64 {
	return v.litres
}

func (v Volume) USGallons() float64 {
	return v.litres / 3.785411784
}

func (v Volume) USLiquidOunces() float64 {
	return v.litres / 0.0295735295625
}

func (v Volume) CubicMeters() float64 {
	return v.litres / 1000
}

func (v Volume) CubicInches() float64 {
	return v.litres / 0.016387064
}

func (v Volume) CubicFeet() float64 {
	return v.litres / 28.316846592
}
