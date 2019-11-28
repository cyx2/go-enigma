package enigma

import "strings"

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Rotor defines a single Enigma rotor
type Rotor struct {
	baseWiring string
	funcWiring string
}

// Plugboard defines a single Enigma plugboard
type Plugboard struct {
	wiring string
}

// InitRotor initializes a single Rotor
func InitRotor(newWiring string) (newRotor *Rotor) {
	newRotor = new(Rotor)
	newRotor.baseWiring, newRotor.funcWiring = newWiring, newWiring
	return newRotor
}

// ReadRotorBaseWiring returns the base (original) wiring of the rotor
func ReadRotorBaseWiring(rotor *Rotor) string {
	return rotor.baseWiring
}

// ReadRotorFuncWiring returns the functional wiring of the rotor
func ReadRotorFuncWiring(rotor *Rotor) string {
	return rotor.funcWiring
}

// ReadRotor takes the readIndex and returns its mapped value
func ReadRotor(rotor *Rotor, readIndex int) (letter string) {
	letter = string(rotor.funcWiring[readIndex])
	return letter
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

// GetPlugboardWiring returns the wiring field in the struct of the plugboard argument
func GetPlugboardWiring(plugboard *Plugboard) (plugboardWiring string) {
	return plugboard.wiring
}

// GetAlphabetIndex returns the index of a letter in the alphabet
func GetAlphabetIndex(letter string) (index int) {
	return strings.Index(alphabet, letter)
}
