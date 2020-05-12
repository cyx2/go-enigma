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
	newRotor.baseWiring, newRotor.offset = newWiring, 0
	return newRotor
}

// InitRotors3 initializes 3 Enigma rotors with the specified rotorWiring arguments
// TODO: Implement variable rotor count
func InitRotors3(rotorWiring1, rotorWiring2, rotorWiring3 string) (rotor1, rotor2, rotor3 *Rotor) {
	rotor1, rotor2, rotor3 = InitRotor(rotorWiring1), InitRotor(rotorWiring2), InitRotor(rotorWiring3)
	return
}

// InitMachine initializes a new 3-rotor Enigma machine
func InitMachine(rotorWiring1, rotorWiring2, rotorWiring3, reflWiring string) (newMachine *Machine) {
	newMachine = new(Machine)
	rotor1, rotor2, rotor3 := InitRotors3(rotorWiring1, rotorWiring2, rotorWiring3)
	rotors := []*Rotor{rotor1, rotor2, rotor3}
	reflector1 := InitReflector(reflWiring)

	newMachine.rotors = rotors
	newMachine.reflector = reflector1
	return newMachine
}

// InitReflector initializes a single Reflector
func InitReflector(newWiring string) (newReflector *Reflector) {
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
