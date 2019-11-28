package enigma

import "strings"

// GetAlphabetIndex returns the index of a letter in the alphabet
func GetAlphabetIndex(letter string) (index int) {
	return strings.Index(alphabet, letter)
}

// GetPlugboardWiring returns the wiring field in the struct of the plugboard argument
func GetPlugboardWiring(plugboard *Plugboard) (plugboardWiring string) {
	return plugboard.wiring
}

// GetReflectorWiring returns the base and refl wiring of a Reflector
func GetReflectorWiring(reflector *Reflector) (baseWiring, reflWiring string) {
	return reflector.baseWiring, reflector.reflWiring
}

// GetRotorWiring returns both the base and functional wiring of a rotor
func GetRotorWiring(rotor *Rotor) (baseWiring, funcWiring string) {
	return rotor.baseWiring, rotor.funcWiring
}
