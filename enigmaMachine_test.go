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

func TestRotorRotateRead(t *testing.T) {
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
}

func TestProcessLetter(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO

	// Base1: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Refl1: JCRBIAPSUXHYWOTNZVQDGLFMKE

	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

	newMachine := enigma.InitMachine(baseWiring1, baseWiring2, baseWiring3, reflBaseWiring)

	testLetter1 := "A"
	processedLetter1 := newMachine.ProcessLetter(testLetter1)
	expProcessedLetter1 := "K"

	testLetter2 := "Z"
	processedLetter2 := newMachine.ProcessLetter(testLetter2)
	expProcessedLetter2 := "E"

	if processedLetter1 != expProcessedLetter1 {
		t.Errorf("Tried to process %v, got %v, want %v", testLetter1, processedLetter1, expProcessedLetter1)
	}

	if processedLetter2 != expProcessedLetter2 {
		t.Errorf("Tried to process %v, got %v, want %v", testLetter2, processedLetter2, expProcessedLetter2)
	}
}
