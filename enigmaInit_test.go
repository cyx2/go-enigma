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
