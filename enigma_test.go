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

func TestReadForward3(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring, baseWiring, baseWiring)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	result1 := enigma.ReadForward3(rotors, "A")
	expectedResult1 := "T"

	result2 := enigma.ReadForward3(rotors, "F")
	expectedResult2 := "F"

	result3 := enigma.ReadForward3(rotors, "P")
	expectedResult3 := "X"

	if result1 != expectedResult1 {
		t.Errorf("ReadForward3 case 1, got %v want %v", result1, expectedResult1)
	}
	if result2 != expectedResult2 {
		t.Errorf("ReadForward3 case 2, got %v want %v", result2, expectedResult2)
	}
	if result3 != expectedResult3 {
		t.Errorf("ReadForward3 case 3, got %v want %v", result3, expectedResult3)
	}
}
