package enigma

// ReadPlugboard takes a readLetter and returns its mapped value
func ReadPlugboard(plugboard *Plugboard, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(plugboard.wiring[readIndex])
	return letter
}

// ReadRotor takes a readLetter and returns its mapped value
func ReadRotor(rotor *Rotor, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(rotor.funcWiring[readIndex])
	return letter
}

// ReadRotorBackwards takes in a readLetter and returns the mapped alphabet value
func ReadRotorBackwards(rotor *Rotor, readLetter string) (letter string) {
	readIndex := GetWiringIndex(rotor.funcWiring, readLetter)
	letter = string(alphabet[readIndex])
	return letter
}

// ReadReflector takes in a readLetter and returns the reflected letter
func ReadReflector(reflector *Reflector, readLetter string) (letter string) {
	readIndex := GetWiringIndex(reflector.baseWiring, readLetter)
	letter = string(reflector.reflWiring[readIndex])
	return letter
}
