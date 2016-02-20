package element

import (
	"strings"
)

type Element struct {
	atomicNumber int
}

func AllElements() []Element {
	var elements = make([]Element, numElements())
	for i := 0; i < len(elements); i++ {
		elements[i] = Element{i + 1}
	}

	return elements
}

func ElementByAtomicNumber(atomicNumber int) Element {
	return Element{atomicNumber}
}

func ElementBySymbol(symbol string) Element {
	for number, element := range allElements {
		if strings.EqualFold(element.symbol, symbol) {
			return Element{number + 1}
		}
	}

	return Element{}
}

func (e Element) getInfo() elementInfo {
	if e.atomicNumber <= 0 || e.atomicNumber > len(allElements) {
		panic("element: invalid element")
	}
	return allElements[e.atomicNumber-1]
}

// Name returns the name of the element.
func (e Element) Name() string {
	return e.getInfo().name
}

// AtomicNumber returns the atomic number of the element.
func (e Element) AtomicNumber() int {
	return e.atomicNumber
}

// MolarMass returns the molar mass (atomic weight) of the element.
func (e Element) MolarMass() float64 {
	return e.getInfo().mass
}

// Group returns the group the element is in.
func (e Element) Group() int {
	return e.getInfo().group
}

// Period returns the period the element is in.
func (e Element) Period() int {
	return e.getInfo().period
}

// Symbol returns the chemical symbol of the element.
func (e Element) Symbol() string {
	return e.getInfo().symbol
}

// IsZero returns whether the element is zero (uninitialized, invalid, or from
// ElementZero)
func (e Element) IsZero() bool {
	if e.atomicNumber == 0 {
		return true
	}
	return false
}
