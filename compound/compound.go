package compound

import (
	"github.com/1lann/beaker/element"
	"strconv"
)

type Compound struct {
	Element    element.Element
	Components []Compound
	Amount     int
}

func (c Compound) String() string {
	str := ""

	if len(c.Components) > 0 {
		if c.Amount > 1 {
			str += strconv.Itoa(c.Amount)
		}

		for _, component := range c.Components {
			if len(component.Components) > 0 {
				str += "("
			}
			componentStr := component.String()
			numberString := ""
			for i := 0; i < len(componentStr); i++ {
				if componentStr[i] >= '0' && componentStr[i] <= '9' {
					numberString += string(componentStr[i])
				} else {
					break
				}
			}

			if len(numberString) > 0 {
				componentStr = componentStr[len(numberString):]
			}

			str += componentStr

			if len(component.Components) > 0 {
				str += ")"
			}

			if len(numberString) > 0 {
				str += numberString
			}
		}

		return str
	}

	str += c.Element.Symbol()

	if c.Amount > 1 {
		str += strconv.Itoa(c.Amount)
	}

	return str
}

func (c Compound) MolarMass() float64 {
	if len(c.Components) == 0 {
		return c.Element.MolarMass() * float64(c.Amount)
	}

	var totalMass float64 = 0

	for _, component := range c.Components {
		totalMass += component.MolarMass()
	}

	return totalMass * float64(c.Amount)
}

func elementInArray(elem element.Element, array []element.Element) bool {
	for _, cmp := range array {
		if cmp == elem {
			return true
		}
	}

	return false
}

func (c Compound) UniqueElements() []element.Element {
	if len(c.Components) == 0 {
		return []element.Element{c.Element}
	}

	var uniqueElements []element.Element
	for _, component := range c.Components {
		results := component.UniqueElements()
		for _, resultElem := range results {
			if !elementInArray(resultElem, uniqueElements) {
				uniqueElements = append(uniqueElements, resultElem)
			}
		}
	}

	return uniqueElements
}

func (c Compound) CountElement(elem element.Element) int {
	if len(c.Components) == 0 {
		if c.Element == elem {
			return c.Amount
		}
		return 0
	}

	var count = 0
	for _, component := range c.Components {
		count += component.CountElement(elem)
	}
	return count
}
