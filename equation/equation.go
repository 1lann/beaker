package equation

import (
	"github.com/1lann/beaker/compound"
)

type Equation struct {
	Reactants []compound.Compound
	Products  []compound.Compound
	balanced  bool
}

func (e *Equation) Balance() {

}
