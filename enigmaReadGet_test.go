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

func TestReadPlugBoard(t *testing.T) {
	// Alphabet:  ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Plugboard: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	plugboard := enigma.InitPlugboard("EKMFLGDQVZNTOWYHXUSPAIBRCJ")
	readLetter1 := enigma.ReadPlugboard(plugboard, "B")
	correctReadLetter1 := "K"
	readLetter2 := enigma.ReadPlugboard(plugboard, "T")
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

	readLetter1 := enigma.ReadRotor(rotor, "O")
	correctReadLetter1 := "Y"
	readLetter2 := enigma.ReadRotor(rotor, "J")
	correctReadLetter2 := "Z"

	if readLetter1 != correctReadLetter1 {
		t.Errorf("ReadRotor(rotor, 'O'), got %v want %v", readLetter1, correctReadLetter1)
	}
	if readLetter2 != correctReadLetter2 {
		t.Errorf("ReadRotor(rotor, 'J'), got %v want %v", readLetter2, correctReadLetter2)
	}
}

func TestReadRotorBackwards(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	baseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotor := enigma.InitRotor(baseWiring)

	readLetter1 := enigma.ReadRotorBackwards(rotor, "X")
	correctReadLetter1 := "Q"
	readLetter2 := enigma.ReadRotorBackwards(rotor, "W")
	correctReadLetter2 := "N"

	if readLetter1 != correctReadLetter1 {
		t.Errorf("ReadRotor(rotor, 'O'), got %v want %v", readLetter1, correctReadLetter1)
	}
	if readLetter2 != correctReadLetter2 {
		t.Errorf("ReadRotor(rotor, 'J'), got %v want %v", readLetter2, correctReadLetter2)
	}
}

func TestReadReflector(t *testing.T) {
	// Base1: EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Refl1: JCRBIAPSUXHYWOTNZVQDGLFMKE
	// Base2: AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Refl2: EOVFYPNZGQCMTWHLBXURISKDJA
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"

	reflector1 := enigma.InitReflector(baseWiring1)
	reflector2 := enigma.InitReflector(baseWiring2)

	ansBase1 := enigma.ReadReflector(reflector1, "G")
	correctAnsBase1 := "A"
	ansBase2 := enigma.ReadReflector(reflector2, "W")
	correctAnsBase2 := "T"

	reflBase1 := enigma.ReadReflector(reflector1, enigma.ReadReflector(reflector1, "G"))
	correctReflBase1 := "G"

	reflBase2 := enigma.ReadReflector(reflector2, enigma.ReadReflector(reflector2, "W"))
	correctReflBase2 := "W"

	if ansBase1 != correctAnsBase1 {
		t.Errorf("AnsBase1 got %v, want %v", ansBase1, correctAnsBase1)
	}
	if ansBase2 != correctAnsBase2 {
		t.Errorf("AnsBase2 got %v, want %v", ansBase2, correctAnsBase2)
	}
	if reflBase1 != correctReflBase1 {
		t.Errorf("ReflBase1 got %v, want %v", reflBase1, correctReflBase1)
	}
	if reflBase2 != correctReflBase2 {
		t.Errorf("ReflBase2 got %v, want %v", reflBase2, correctReflBase2)
	}
}
