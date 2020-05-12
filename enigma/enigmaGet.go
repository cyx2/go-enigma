package enigma

import "strings"

// GetAlphabetIndex returns the index of a letter in the alphabet
func GetAlphabetIndex(letter string) (index int) {
	return strings.Index(alphabet, letter)
}

// GetAlphabetLetter returns the letter at the argument index in the alphabet
func GetAlphabetLetter(index int) (letter string) {
	return string(alphabet[index])
}

// GetWiringIndex returns the index of a letter in the wiring argument
func GetWiringIndex(wiring, letter string) (index int) {
	return strings.Index(wiring, letter)
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
func GetRotorWiring(rotor *Rotor) (wiring string) {
	return rotor.baseWiring
}

// GetRotorOffset returns the offset of a rotor
func GetRotorOffset(rotor *Rotor) (offset int) {
	return rotor.offset
}

// GetMachineRotorWiring returns the base wiring of a rotor in a machine
func GetMachineRotorWiring(machine *Machine, rotorIndex int) (baseWiring string) {
	rotors := machine.rotors
	baseWiring = GetRotorWiring(rotors[rotorIndex])
	return baseWiring
}

// GetMachineRotorOffset returns the offset of a rotor in a machine
func GetMachineRotorOffset(machine *Machine, rotorIndex int) (rotorOffset int) {
	rotors := machine.rotors
	rotorOffset = GetRotorOffset(rotors[rotorIndex])
	return rotorOffset
}
