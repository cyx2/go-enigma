package main

import (
	"go-enigma/enigma"
	"testing"
)

func TestRotorRotate(t *testing.T) {
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)
	rotor.Rotate()
	_, newFuncWiring := enigma.GetRotorWiring(rotor)
	expectedFuncWiring := "KMFLGDQVZNTOWYHXUSPAIBRCJE"

	if newFuncWiring != expectedFuncWiring {
		t.Errorf("InitRotor('EKMFLGDQVZNTOWYHXUSPAIBRCJ'), got %v want %v", newFuncWiring, expectedFuncWiring)
	}
}

func TestRotorDoubleRotate(t *testing.T) {
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)
	rotor.Rotate()
	rotor.Rotate()
	_, newFuncWiring := enigma.GetRotorWiring(rotor)
	expectedFuncWiring := "MFLGDQVZNTOWYHXUSPAIBRCJEK"

	if newFuncWiring != expectedFuncWiring {
		t.Errorf("InitRotor('EKMFLGDQVZNTOWYHXUSPAIBRCJ'), got %v want %v", newFuncWiring, expectedFuncWiring)
	}
}
