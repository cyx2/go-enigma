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
