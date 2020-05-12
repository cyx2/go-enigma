package enigma

// Rotate function rotates a roter one notch
func (rotor *Rotor) Rotate() {
	currentOffset := GetRotorOffset(rotor)

	// Prevent the offset from going beyond the alphabet
	if currentOffset < 25 {
		rotor.offset++
	} else {
		rotor.offset = 0
	}
}

// ProcessLetter will process a single enigma letter
func (machine *Machine) ProcessLetter(letter string) (processedLetter string) {
	machineRotors := machine.rotors
	machineReflector := machine.reflector

	processForward := ReadForward(machineRotors, letter)
	processReflector := ReadReflector(machineReflector, processForward)
	processBackward := ReadBackward(machineRotors, processReflector)

	processedLetter = processBackward

	return processedLetter
}
