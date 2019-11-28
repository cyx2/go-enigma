package enigma

// ReadPlugboard takes the readIndex and returns its mapped value
func ReadPlugboard(plugboard *Plugboard, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(plugboard.wiring[readIndex])
	return letter
}

// ReadRotor takes the readIndex and returns its mapped value
func ReadRotor(rotor *Rotor, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(rotor.funcWiring[readIndex])
	return letter
}

// ReadRotorBackwards takes in a readIndex and returns the mapped alphabet value
func ReadRotorBackwards(rotor *Rotor, readLetter string) (letter string) {
	readIndex := GetWiringIndex(rotor.funcWiring, readLetter)
	letter = string(alphabet[readIndex])
	return letter
}
