package main

import (
	"go-enigma/enigma"
	"testing"
)

func TestSetRotorOffset(t *testing.T) {
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	newRotor1 := enigma.InitRotor(baseWiring1)

	initOffset := enigma.GetRotorOffset(newRotor1)
	expInitOffset := 0

	if initOffset != expInitOffset {
		t.Errorf("Initialized new rotor, got offset %v, want %v", initOffset, expInitOffset)
	}

	testRotorOffset1 := 5
	newRotor1.SetRotorOffset(testRotorOffset1)
	updatedRotorOffset1 := enigma.GetRotorOffset(newRotor1)

	if updatedRotorOffset1 != testRotorOffset1 {
		t.Errorf("Set new offset, got %v, want %v", updatedRotorOffset1, testRotorOffset1)
	}

	testRotorOffset2 := 100
	newRotor1.SetRotorOffset(testRotorOffset2)
	updatedRotorOffset2 := enigma.GetRotorOffset(newRotor1)

	if updatedRotorOffset2 != testRotorOffset1 {
		t.Errorf("Shouldn't have been able to set new offset of %v, want it to stay as %v", updatedRotorOffset2, testRotorOffset1)
	}
}
