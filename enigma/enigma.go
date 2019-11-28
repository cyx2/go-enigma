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

// Reflector defines a single Enigma reflector
type Reflector struct {
	baseWiring string
	reflWiring string
}

// ReadRotor takes the readIndex and returns its mapped value
func ReadRotor(rotor *Rotor, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(rotor.funcWiring[readIndex])
	return letter
}

// ReadPlugboard takes the readIndex and returns its mapped value
func ReadPlugboard(plugboard *Plugboard, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(plugboard.wiring[readIndex])
	return letter
}

// Rotate function rotates a roter one notch
func (rotor *Rotor) Rotate() {
	_, currentFuncWiring := GetRotorWiring(rotor)
	firstLetter := string(currentFuncWiring[0])

	newFuncWiring := strings.Join([]string{currentFuncWiring, firstLetter}, "")
	newFuncWiringRunes := []rune(newFuncWiring)[1 : len(alphabet)+1]
	rotor.funcWiring = string(newFuncWiringRunes)
}
