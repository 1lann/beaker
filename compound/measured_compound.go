package compound

import (
	"github.com/1lann/beaker/element"
	"github.com/1lann/beaker/physics"
)

type MeasuredCompound struct {
	Compound
	moles float64
}

type GasCondition struct {
	Temperature physics.Temperature
	Pressure    physics.Pressure
	Volume      physics.Volume
}

func MolesByGasCondition(condition GasCondition) float64 {
	return (condition.Pressure.KiloPascals() * condition.Volume.Litres()) /
		(physics.R * condition.Temperature.Kelvin())
}

func (c Compound) MeasureByMass(mass physics.Mass) MeasuredCompound {
	return MeasuredCompound{
		Compound: c,
		moles:    mass.Grams() / c.MolarMass(),
	}
}

func (c Compound) MeasureByMoles(moles float64) MeasuredCompound {
	return MeasuredCompound{
		Compound: c,
		moles:    moles,
	}
}

func (c Compound) MeasureByGasCondition(
	condition GasCondition) MeasuredCompound {
	return MeasuredCompound{
		Compound: c,
		moles:    MolesByGasCondition(condition),
	}
}

func (m MeasuredCompound) Mass() physics.Mass {
	return physics.Grams(m.moles * m.MolarMass())
}

func (m MeasuredCompound) Moles() float64 {
	return m.moles
}

func (m MeasuredCompound) MolesMultAmount() float64 {
	return m.moles * float64(m.Amount)
}

func (m MeasuredCompound) MeasuredComponents() []MeasuredCompound {
	var measuredComponents = make([]MeasuredCompound, len(m.Components))
	for i, compound := range m.Components {
		measuredComponents[i] = compound.MeasureByMoles(m.moles *
			float64(m.Amount))
	}
	return measuredComponents
}

func (m MeasuredCompound) TotalMolesElement(elem element.Element) float64 {
	return m.moles * float64(m.CountElement(elem))
}

func (m MeasuredCompound) TotalMassElement(elem element.Element) float64 {
	return m.TotalMolesElement(elem) * elem.MolarMass()
}
