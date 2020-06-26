package main

import (
	"go-enigma/enigma"
	"testing"
)

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
	rotorWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(rotorWiring)
	newRotorBaseWiring := enigma.GetRotorWiring(rotor)
	newRotorOffset := enigma.GetRotorOffset(rotor)
	expNewRotorOffset := 0

	if newRotorBaseWiring != rotorWiring {
		t.Errorf("Got rotor base wiring of %v, want %v", newRotorBaseWiring, rotorWiring)
	}
	if newRotorOffset != expNewRotorOffset {
		t.Errorf("Got rotor offset wiring of %v, want %v", newRotorOffset, expNewRotorOffset)
	}
}

func TestInitReflector(t *testing.T) {
	// Base Wiring: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Refl Wiring: JCRBIAPSUXHYWOTNZVQDGLFMKE
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	reflector := enigma.InitReflector(reflBaseWiring)
	baseWiring := enigma.GetReflectorWiring(reflector)

	if baseWiring != reflBaseWiring {
		t.Errorf("Got reflector base wiring of %v, want %v", baseWiring, reflBaseWiring)
	}
}
