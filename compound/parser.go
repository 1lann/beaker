package compound

import (
	"github.com/1lann/beaker/element"
	"strconv"
)

// type Compound struct {
// 	Element    element.Element
// 	Components []Compound
// 	Amount     int
// }

type ParseError struct {
	location     int
	errorMessage string
}

func (p *ParseError) Error() string {
	return "compound: parse error: " + p.errorMessage
}

func (p *ParseError) Message() string {
	return p.errorMessage
}

func (p *ParseError) Location() int {
	return p.location
}

func ParseString(str string) (Compound, error) {
	var layers = []Compound{Compound{Amount: 1}}
	var layerNum = 0

	for i := 0; i < len(str); i++ {
		currentLetter := string(str[i])

		if currentLetter == "(" {
			layerNum++
			layers = append(layers, Compound{Amount: 1})
		} else if currentLetter == ")" {
			if len(layers[layerNum].Components) == 0 {
				return Compound{}, &ParseError{
					location:     i - 1,
					errorMessage: "Expected content inside brackets",
				}
			}

			layerNum--
			if layerNum < 0 {
				return Compound{}, &ParseError{
					location:     i,
					errorMessage: "Unexpected closing bracket \")\"",
				}
			}

			layers[layerNum].Components = append(layers[layerNum].Components,
				layers[layerNum+1])
			layers = layers[:layerNum+1]
		} else if currentLetter[0] >= '0' && currentLetter[0] <= '9' {
			// Consume number
			var numberString = currentLetter
			for {
				i++
				if i < len(str) && str[i] >= '0' && str[i] <= '9' {
					numberString += string(str[i])
				} else {
					i--
					break
				}
			}

			number, err := strconv.Atoi(numberString)
			if err != nil {
				return Compound{}, err
			}

			if i-len(numberString)+1 == 0 {
				layers[layerNum].Amount = number
			} else {
				layers[layerNum].Components[len(
					layers[layerNum].Components)-1].Amount = number
			}
		} else if currentLetter[0] >= 'A' && currentLetter[0] <= 'Z' {
			if i+1 < len(str) && str[i+1] >= 'a' && str[i+1] <= 'z' {
				// Is a two letter element
				symbol := currentLetter + string(str[i+1])
				elem := element.ElementBySymbol(symbol)
				if elem.IsZero() {
					return Compound{}, &ParseError{
						location:     i,
						errorMessage: "Non-existent element: " + symbol,
					}
				}

				i++

				layers[layerNum].Components = append(
					layers[layerNum].Components,
					Compound{
						elem,
						[]Compound{},
						1,
					})
			} else {
				// Is a one letter element
				elem := element.ElementBySymbol(currentLetter)
				if elem.IsZero() {
					return Compound{}, &ParseError{
						location:     i,
						errorMessage: "Non-existent element: " + currentLetter,
					}
				}

				layers[layerNum].Components = append(
					layers[layerNum].Components,
					Compound{
						elem,
						[]Compound{},
						1,
					})
			}
		} else {
			return Compound{}, &ParseError{
				location:     i,
				errorMessage: "Unexpected character: " + currentLetter,
			}
		}
	}

	if layerNum > 0 {
		return Compound{}, &ParseError{
			location:     0,
			errorMessage: "Missing closing bracket \")\"",
		}
	}

	return layers[0], nil
}

func MustParseString(str string) Compound {
	cmp, err := ParseString(str)
	if err != nil {
		panic(err)
	}

	return cmp
}
