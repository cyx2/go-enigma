package enigma

// InitPlugboard initializes a single Plugboard
func InitPlugboard(newWiring string) (newPlugboard *Plugboard) {
	newPlugboard = new(Plugboard)
	newPlugboard.wiring = newWiring
	return newPlugboard
}

// InitRotor initializes a single Rotor
func InitRotor(newWiring string) (newRotor *Rotor) {
	newRotor = new(Rotor)
	newRotor.baseWiring, newRotor.funcWiring = newWiring, newWiring
	return newRotor
}

// InitReflector initializes a single Reflector
func InitReflector(newWiring string) (newReflector *Reflector) {
	// newRotor = new(Rotor)
	// newRotor.baseWiring, newRotor.funcWiring = newWiring, newWiring
	// return newRotor

	newReflector = new(Reflector)
	newReflector.baseWiring = newWiring
	newReflector.reflWiring = reverse(newWiring)
	return newReflector
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
