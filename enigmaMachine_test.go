package main

import (
	"go-enigma/enigma"
	"testing"
)

func TestRotorRotate(t *testing.T) {
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)
	rotor.Rotate()
	offset := enigma.GetRotorOffset(rotor)
	expOffset := 1

	if offset != expOffset {
		t.Errorf("Rotate(rotor) single, got offset %v want %v", offset, expOffset)
	}

	rotor.Rotate()
	rotor.Rotate()

	offset = enigma.GetRotorOffset(rotor)
	expOffset = 3
	if offset != expOffset {
		t.Errorf("Rotate(rotor) triple, got offset %v want %v", offset, expOffset)
	}
}

func TestRotorRotateOffset(t *testing.T) {
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)

	readLetter := enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, 0))
	expReadLetter := "E"

	if readLetter != expReadLetter {
		t.Errorf("No rotation, rotor offset %v, read %v want %v", enigma.GetRotorOffset(rotor), readLetter, expReadLetter)
	}

	rotor.Rotate()

	readLetter = enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, 0))
	expReadLetter = "K"
	if readLetter != expReadLetter {
		t.Errorf("Single rotation, rotor offset %v, read %v want %v", enigma.GetRotorOffset(rotor), readLetter, expReadLetter)
	}

	rotor.Rotate()
	rotor.Rotate()

	readLetter = enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, 0))
	expReadLetter = "F"
	if readLetter != expReadLetter {
		t.Errorf("Single rotation, rotor offset %v, read %v want %v", enigma.GetRotorOffset(rotor), readLetter, expReadLetter)
	}
}

func TestRotorRotateReadForward(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ

	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	singleRotor := enigma.InitRotor(baseWiring)
	rotors := []*enigma.Rotor{singleRotor}

	testLetter := "A"
	forwardRead1 := enigma.ReadForward(rotors, testLetter)
	expForwardRead1 := "E"

	if forwardRead1 != expForwardRead1 {
		t.Errorf("Tried to %v read with offset %v, got %v, want %v", testLetter, enigma.GetRotorOffset(singleRotor), forwardRead1, expForwardRead1)
	}

	singleRotor.Rotate()

	forwardRead2 := enigma.ReadForward(rotors, testLetter)
	expForwardRead2 := "K"

	if forwardRead2 != expForwardRead2 {
		t.Errorf("Tried to %v read with offset %v, got %v, want %v", testLetter, enigma.GetRotorOffset(singleRotor), forwardRead2, expForwardRead2)
	}
}

func TestRotorRotateReadBackward(t *testing.T) {
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	singleRotor := enigma.InitRotor(baseWiring)
	rotors := []*enigma.Rotor{singleRotor}

	testLetter := "H"
	forwardRead1 := enigma.ReadBackward(rotors, testLetter)
	expForwardRead1 := "P"

	if forwardRead1 != expForwardRead1 {
		t.Errorf("Tried to %v read with offset %v, got %v, want %v", testLetter, enigma.GetRotorOffset(singleRotor), forwardRead1, expForwardRead1)
	}

	singleRotor.Rotate()

	forwardRead2 := enigma.ReadBackward(rotors, testLetter)
	expForwardRead2 := "O"

	if forwardRead2 != expForwardRead2 {
		t.Errorf("Tried to %v read with offset %v, got %v, want %v", testLetter, enigma.GetRotorOffset(singleRotor), forwardRead2, expForwardRead2)
	}
}

func TestRotorEndBound(t *testing.T) {
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)
	rotor.SetRotorOffset(24)

	readIndex := 0

	readLetter1 := enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, readIndex))
	expReadLetter1 := "C"

	if readLetter1 != expReadLetter1 {
		t.Errorf("Rotor with offset %v got letter %v, want %v", enigma.GetRotorOffset(rotor), readLetter1, expReadLetter1)
	}

	rotor.Rotate()

	readLetter2 := enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, readIndex))
	expReadLetter2 := "J"

	if readLetter2 != expReadLetter2 {
		t.Errorf("Rotor with offset %v got letter %v, want %v", enigma.GetRotorOffset(rotor), readLetter2, expReadLetter2)
	}

	// Rotate past end bound, expect offset to return to 0
	rotor.Rotate()

	readLetter3 := enigma.GetAlphabetLetter(enigma.ReadRotor(rotor, readIndex))
	expReadLetter3 := "E"

	if readLetter3 != expReadLetter3 {
		t.Errorf("Rotor with offset %v got letter %v, want %v", enigma.GetRotorOffset(rotor), readLetter3, expReadLetter3)
	}
}

func TestMachineInit(t *testing.T) {
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

	newMachine := enigma.InitMachine(baseWiring1, baseWiring2, baseWiring3, reflBaseWiring)

	rotor1Index := 0
	newMachineRotor1Wiring := enigma.GetMachineRotorWiring(newMachine, rotor1Index)

	if newMachineRotor1Wiring != baseWiring1 {
		t.Errorf("For rotor %v got wiring of %v, want %v", rotor1Index, newMachineRotor1Wiring, baseWiring1)
	}

	rotor3Index := 2
	rotor3Offset := 0
	newMachineRotor3Wiring := enigma.GetMachineRotorWiring(newMachine, rotor3Index)
	newMachineRotor3Offset := enigma.GetMachineRotorOffset(newMachine, rotor3Index)

	if newMachineRotor3Wiring != baseWiring3 {
		t.Errorf("For rotor %v got wiring of %v, want %v", rotor3Index, newMachineRotor3Wiring, baseWiring3)
	}
	if newMachineRotor3Offset != rotor3Offset {
		t.Errorf("For rotor %v got offset of %v, want %v", rotor3Index, newMachineRotor3Offset, rotor3Offset)
	}

	numLettersProcessed := enigma.GetMachineNumLettersProcessed(newMachine)
	expNumLettersProcessed := 0

	if numLettersProcessed != expNumLettersProcessed {
		t.Errorf("Machine processed no letters, got %v, want %v", numLettersProcessed, expNumLettersProcessed)
	}
}

func TestProcessLetter(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ

	// Alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Refl1:    YRUHQSLDPXNGOKMIEBFZCWVJAT

	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	reflBaseWiring := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	newMachine := enigma.InitMachine(baseWiring1, baseWiring2, baseWiring3, reflBaseWiring)

	testLetter1 := "A"
	processedLetter1 := newMachine.ProcessLetter(testLetter1)
	expProcessedLetter1 := "U"

	testLetter2 := "E"
	processedLetter2 := newMachine.ProcessLetter(testLetter2)
	expProcessedLetter2 := "B"

	if processedLetter1 != expProcessedLetter1 {
		t.Errorf("Tried to process %v, got %v, want %v", testLetter1, processedLetter1, expProcessedLetter1)
	}

	if processedLetter2 != expProcessedLetter2 {
		t.Errorf("Tried to process %v, got %v, want %v", testLetter2, processedLetter2, expProcessedLetter2)
	}
}

// func TestProcessString(t *testing.T) {
// 	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
// 	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
// 	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
// 	reflBaseWiring := "YRUHQSLDPXNGOKMIEBFZCWVJAT"

// 	newMachine := enigma.InitMachine(baseWiring1, baseWiring2, baseWiring3, reflBaseWiring)

// 	testString := "AAA"
// 	processedString := newMachine.ProcessString(testString)
// 	expProcessedString := "DHL"

// 	if processedString != expProcessedString {
// 		t.Errorf("Tried to process string %v, got %v, want %v", testString, processedString, expProcessedString)
// 	}

// 	numLettersProcessed := enigma.GetMachineNumLettersProcessed(newMachine)
// 	expNumLettersProcessed := 3

// 	if numLettersProcessed != expNumLettersProcessed {
// 		t.Errorf("Machine processed letters, got %v, want %v", numLettersProcessed, expNumLettersProcessed)
// 	}

// }
