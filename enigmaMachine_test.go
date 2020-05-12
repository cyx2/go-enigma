package main

import (
	"go-enigma/enigma"
	"testing"
)

func TestMachineInit(t *testing.T) {
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
