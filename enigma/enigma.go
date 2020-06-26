package enigma

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Rotor defines a single Enigma rotor
type Rotor struct {
	baseWiring string
	offset     int
}

// Plugboard defines a single Enigma plugboard
type Plugboard struct {
	wiring string
}

// Reflector defines a single Enigma reflector
type Reflector struct {
	baseWiring string
}

// Machine defines a single Enigma machine
type Machine struct {
	rotors              []*Rotor
	reflector           *Reflector
	numLettersProcessed int
}
