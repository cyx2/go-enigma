package main

import (
	"go-enigma/enigma"
	"testing"
)

func TestGetAlphabetIndex(t *testing.T) {
	ans1 := enigma.GetAlphabetIndex("B")
	correctAns1 := 1
	ans2 := enigma.GetAlphabetIndex("Z")
	correctAns2 := 25
	if ans1 != correctAns1 {
		t.Errorf("GetAlphabetIndex('B'), got %v want %v", ans1, correctAns1)
	}
	if ans2 != correctAns2 {
		t.Errorf("GetAlphabetIndex('B'), got %v want %v", ans2, correctAns2)
	}
}

func TestInitPlugboard(t *testing.T) {
	// Plugboard: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	plugboardWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	plugboard := enigma.InitPlugboard(plugboardWiring)
	ans := enigma.GetPlugboardWiring(plugboard)

	if ans != plugboardWiring {
		t.Errorf("Got wiring of %v, want %v", ans, plugboardWiring)
	}
}

func TestInitRotor(t *testing.T) {
	// Base Wiring: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Func Wiring: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	rotorWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(rotorWiring)
	newRotorBaseWiring, newRotorFuncWiring := enigma.GetRotorWiring(rotor)

	if newRotorBaseWiring != rotorWiring {
		t.Errorf("Got rotor base wiring of %v, want %v", newRotorBaseWiring, rotorWiring)
	}
	if newRotorFuncWiring != rotorWiring {
		t.Errorf("Got rotor func wiring of %v, want %v", newRotorFuncWiring, rotorWiring)
	}
}

func TestInitReflector(t *testing.T) {
	// Base Wiring: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Refl Wiring: JCRBIAPSUXHYWOTNZVQDGLFMKE
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	reflector := enigma.InitReflector(reflBaseWiring)
	expReflWiring := "JCRBIAPSUXHYWOTNZVQDGLFMKE"
	baseWiring, reflWiring := enigma.GetReflectorWiring(reflector)

	if baseWiring != reflBaseWiring {
		t.Errorf("Got reflector base wiring of %v, want %v", baseWiring, reflBaseWiring)
	}
	if reflWiring != expReflWiring {
		t.Errorf("Got reflector refl wiring of %v, want %v", reflWiring, expReflWiring)
	}
}

func TestReadPlugBoard(t *testing.T) {
	// Alphabet:  ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Plugboard: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	plugboard := enigma.InitPlugboard("EKMFLGDQVZNTOWYHXUSPAIBRCJ")
	readLetter1 := enigma.ReadPlugboard(plugboard, enigma.GetAlphabetIndex("B"))
	correctReadLetter1 := "K"
	readLetter2 := enigma.ReadPlugboard(plugboard, enigma.GetAlphabetIndex("T"))
	correctReadLetter2 := "P"

	if readLetter1 != correctReadLetter1 {
		t.Errorf("ReadPlugboard(plugboard, enigma.GetAlphabetIndex('B')), got %v want %v", readLetter1, correctReadLetter1)
	}
	if readLetter2 != correctReadLetter2 {
		t.Errorf("ReadPlugboard(plugboard, enigma.GetAlphabetIndex('T')), got %v want %v", readLetter2, correctReadLetter2)
	}
}

func TestReadRotor(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)

	readLetter1 := enigma.ReadRotor(rotor, enigma.GetAlphabetIndex("O"))
	correctReadLetter1 := "Y"
	readLetter2 := enigma.ReadRotor(rotor, enigma.GetAlphabetIndex("J"))
	correctReadLetter2 := "Z"

	if readLetter1 != correctReadLetter1 {
		t.Errorf("ReadRotor(rotor, enigma.GetAlphabetIndex('O')), got %v want %v", readLetter1, correctReadLetter1)
	}
	if readLetter2 != correctReadLetter2 {
		t.Errorf("ReadRotor(rotor, enigma.GetAlphabetIndex('J')), got %v want %v", readLetter2, correctReadLetter2)
	}
}

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
