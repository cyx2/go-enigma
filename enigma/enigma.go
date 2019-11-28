package enigma

import "strings"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Rotor defines a single Enigma rotor
type Rotor struct {
	wiring string
	index  uint8
}

// Plugboard defines a single Enigma plugboard
type Plugboard struct {
	wiring string
}

// InitRotor initializes a single Rotor
func InitRotor(newWiring string) (newRotor *Rotor) {
	newRotor = new(Rotor)
	newRotor.wiring = newWiring
	newRotor.index = 1
	return newRotor
}

// InitPlugboard initializes a single Plugboard
func InitPlugboard(newWiring string) (newPlugboard *Plugboard) {
	newPlugboard = new(Plugboard)
	newPlugboard.wiring = newWiring
	return newPlugboard
}

// ReadPlugboard takes the readIndex and returns its mapped value
func ReadPlugboard(plugboard *Plugboard, readIndex int) (letter string) {
	letter = string(plugboard.wiring[readIndex])
	return letter
}

// GetAlphabetIndex returns the index of a letter in the alphabet
func GetAlphabetIndex(letter string) (index int) {
	return strings.Index(alphabet, letter)
}
