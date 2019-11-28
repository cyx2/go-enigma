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

func TestRotorRoundTrip(t *testing.T) {
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

	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring1, baseWiring2, baseWiring3)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}
	reflector := enigma.InitReflector(reflBaseWiring)

	forwardRead := enigma.ReadForward(rotors, "A")
	expForwardRead := "G"

	if forwardRead != expForwardRead {
		t.Errorf("Forward read failed, got %v, want %v", forwardRead, expForwardRead)
	}

	reflectorRead := enigma.ReadReflector(reflector, forwardRead)
	expReflectorRead := "A"

	if reflectorRead != expReflectorRead {
		t.Errorf("Reflector read failed, got %v, want %v", reflectorRead, expReflectorRead)
	}

	backwardRead := enigma.ReadBackward(rotors, reflectorRead)
	expBackwardRead := "K"

	if backwardRead != expBackwardRead {
		t.Errorf("Backward read failed, got %v, want %v", backwardRead, expBackwardRead)
	}
}
