package enigma

// ReadPlugboard takes a readLetter and returns its mapped value
func ReadPlugboard(plugboard *Plugboard, readLetter string) (letter string) {
	readIndex := GetAlphabetIndex(readLetter)
	letter = string(plugboard.wiring[readIndex])
	return letter
}

// ReadRotor takes a readIndex and returns its mapped index
func ReadRotor(rotor *Rotor, readIndex int) (retIndex int) {
	baseIndex := readIndex + GetRotorOffset(rotor)
	baseIndex = offsetAdjust(baseIndex)

	retLetter := string(rotor.baseWiring[baseIndex])
	retIndex = GetAlphabetIndex(retLetter)
	return retIndex
}

// ReadRotorBackwards takes a readIndex and returns its mapped alphabet index
func ReadRotorBackwards(rotor *Rotor, readIndex int) (retIndex int) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	baseIndex := readIndex + GetRotorOffset(rotor)
	baseIndex = offsetAdjust(baseIndex)

	retLetter := GetAlphabetLetter(baseIndex)
	retIndex = GetAlphabetIndex(retLetter)
	return retIndex
}

// ReadReflector takes in a readLetter and returns the reflected letter
func ReadReflector(reflector *Reflector, readLetter string) (letter string) {
	readIndex := GetWiringIndex(reflector.baseWiring, readLetter)
	letter = string(reflector.reflWiring[readIndex])
	return letter
}

// ReadForward forward encrypts letter using 3 rotors in sequence
func ReadForward(rotors []*Rotor, letter string) (encryptedLetter string) {
	encryptedLetter = letter
	encryptedIndex := GetAlphabetIndex(encryptedLetter)
	for i := range rotors {
		encryptedIndex = ReadRotor(rotors[i], encryptedIndex)
	}
	encryptedLetter = GetAlphabetLetter(encryptedIndex)
	return encryptedLetter
}

// ReadBackward backward encrypts letter using 3 rotors in sequence
func ReadBackward(rotors []*Rotor, letter string) (encryptedLetter string) {
	// Reverse rotors to read them backwards
	for i, j := 0, len(rotors)-1; i < j; i, j = i+1, j-1 {
		rotors[i], rotors[j] = rotors[j], rotors[i]
	}

	encryptedLetter = letter
	var encryptedIndex int
	for i := range rotors {
		wiringIndex := GetWiringIndex(rotors[i].baseWiring, encryptedLetter)
		encryptedIndex = ReadRotorBackwards(rotors[i], wiringIndex)
		encryptedLetter = GetAlphabetLetter(encryptedIndex)
	}

	// Return rotor order back to initial sequence
	for i, j := 0, len(rotors)-1; i < j; i, j = i+1, j-1 {
		rotors[i], rotors[j] = rotors[j], rotors[i]
	}

	encryptedLetter = GetAlphabetLetter(encryptedIndex)
	return encryptedLetter
}
