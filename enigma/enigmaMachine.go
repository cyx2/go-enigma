package enigma

import (
	"math"
	"strings"
)

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

// ProcessLetter will process a single letter
func (machine *Machine) ProcessLetter(letter string) (processedLetter string) {
	machineRotors := machine.rotors
	machineReflector := machine.reflector

	processForward := ReadForward(machineRotors, letter)
	processReflector := ReadReflector(machineReflector, processForward)
	processBackward := ReadBackward(machineRotors, processReflector)

	processedLetter = processBackward

	return processedLetter
}

// ProcessString will process a string of many letters
func (machine *Machine) ProcessString(manyLetters string) (output string) {
	inputLetters := []rune(manyLetters)
	processedLetters := make([]string, len(inputLetters))

	for index, letter := range inputLetters {
		machine.rotors[2].Rotate()
		if math.Mod(float64(machine.numLettersProcessed), 25) == 0 {
			machine.rotors[1].Rotate()
		}
		if math.Mod(float64(machine.numLettersProcessed), 50) == 0 {
			machine.rotors[0].Rotate()
		}

		letterToBeProcessed := string(letter)
		processedLetters[index] = machine.ProcessLetter(letterToBeProcessed)
		machine.numLettersProcessed++
	}

	output = strings.Join(processedLetters, "")
	return output
}
