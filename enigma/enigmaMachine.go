package enigma

// Rotate function rotates a roter one notch
func (rotor *Rotor) Rotate() {
	rotor.offset++
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
