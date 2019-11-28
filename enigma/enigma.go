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

// Rotate function rotates a roter one notch
func (rotor *Rotor) Rotate() {
	_, currentFuncWiring := GetRotorWiring(rotor)
	firstLetter := string(currentFuncWiring[0])

	newFuncWiring := strings.Join([]string{currentFuncWiring, firstLetter}, "")
	newFuncWiringRunes := []rune(newFuncWiring)[1 : len(alphabet)+1]
	rotor.funcWiring = string(newFuncWiringRunes)
}

// InitRotors3 initializes 3 Enigma rotors with the specified rotorWiring arguments
func InitRotors3(rotorWiring1, rotorWiring2, rotorWiring3 string) (rotor1, rotor2, rotor3 *Rotor) {
	rotor1, rotor2, rotor3 = InitRotor(rotorWiring1), InitRotor(rotorWiring2), InitRotor(rotorWiring3)
	return
}
