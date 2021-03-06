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

	readIndex1 := enigma.ReadRotor(rotor, enigma.GetAlphabetIndex("O"))
	readLetter1 := enigma.GetAlphabetLetter(readIndex1)
	correctReadLetter1 := "Y"
	readIndex2 := enigma.ReadRotor(rotor, enigma.GetAlphabetIndex("J"))
	readLetter2 := enigma.GetAlphabetLetter(readIndex2)
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

	readIndex1 := enigma.ReadRotorBackwards(rotor, enigma.GetWiringIndex(baseWiring, "X"))
	readLetter1 := enigma.GetAlphabetLetter(readIndex1)
	correctReadLetter1 := "Q"
	readIndex2 := enigma.ReadRotorBackwards(rotor, enigma.GetWiringIndex(baseWiring, "W"))
	readLetter2 := enigma.GetAlphabetLetter(readIndex2)
	correctReadLetter2 := "N"

	if readLetter1 != correctReadLetter1 {
		t.Errorf("ReadRotor(rotor, 'X'), got %v want %v", readLetter1, correctReadLetter1)
	}
	if readLetter2 != correctReadLetter2 {
		t.Errorf("ReadRotor(rotor, 'W'), got %v want %v", readLetter2, correctReadLetter2)
	}
}

func TestReadReflector(t *testing.T) {
	// Alpha: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Base1: EKMFLGDQVZNTOWYHXUSPAIBRCJ

	// Alpha: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Base2: AJDKSIRUXBLHWTMCQGZNPYFVOE
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"

	reflector1 := enigma.InitReflector(baseWiring1)
	reflector2 := enigma.InitReflector(baseWiring2)

	ansBase1 := enigma.ReadReflector(reflector1, "G")
	correctAnsBase1 := "D"
	ansBase2 := enigma.ReadReflector(reflector2, "W")
	correctAnsBase2 := "F"

	if ansBase1 != correctAnsBase1 {
		t.Errorf("AnsBase1 got %v, want %v", ansBase1, correctAnsBase1)
	}
	if ansBase2 != correctAnsBase2 {
		t.Errorf("AnsBase2 got %v, want %v", ansBase2, correctAnsBase2)
	}
}

func TestReadForward(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring1, baseWiring2, baseWiring3)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	result1 := enigma.ReadForward(rotors, "A")
	expectedResult1 := "Z"

	result2 := enigma.ReadForward(rotors, "F")
	expectedResult2 := "Q"

	result3 := enigma.ReadForward(rotors, "S")
	expectedResult3 := "U"

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

func TestReadBackwardSimple1(t *testing.T) {
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	rotors := []*enigma.Rotor{enigma.InitRotor(baseWiring1)}

	result1 := enigma.ReadBackward(rotors, "E")
	expectedResult1 := "A"

	result2 := enigma.ReadBackward(rotors, "J")
	expectedResult2 := "Z"

	result3 := enigma.ReadBackward(rotors, "T")
	expectedResult3 := "L"

	if result1 != expectedResult1 {
		t.Errorf("ReadBackward case 1, got %v want %v", result1, expectedResult1)
	}
	if result2 != expectedResult2 {
		t.Errorf("ReadBackward case 2, got %v want %v", result2, expectedResult2)
	}
	if result3 != expectedResult3 {
		t.Errorf("ReadBackward case 3, got %v want %v", result3, expectedResult3)
	}
}

func TestReadBackwardSimple2(t *testing.T) {
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	rotor1, rotor2 := enigma.InitRotor(baseWiring1), enigma.InitRotor(baseWiring2)
	rotors := []*enigma.Rotor{rotor1, rotor2}

	result1 := enigma.ReadBackward(rotors, "A")
	expectedResult1 := "H"

	result2 := enigma.ReadBackward(rotors, "E")
	expectedResult2 := "A"

	result3 := enigma.ReadBackward(rotors, "D")
	expectedResult3 := "R"

	if result1 != expectedResult1 {
		t.Errorf("ReadBackward case 1, got %v want %v", result1, expectedResult1)
	}
	if result2 != expectedResult2 {
		t.Errorf("ReadBackward case 2, got %v want %v", result2, expectedResult2)
	}
	if result3 != expectedResult3 {
		t.Errorf("ReadBackward case 3, got %v want %v", result3, expectedResult3)
	}
}

func TestReadBackward(t *testing.T) {
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring1, baseWiring2, baseWiring3)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}

	result1 := enigma.ReadBackward(rotors, "A")
	expectedResult1 := "D"

	result2 := enigma.ReadBackward(rotors, "F")
	expectedResult2 := "G"

	result3 := enigma.ReadBackward(rotors, "P")
	expectedResult3 := "N"

	if result1 != expectedResult1 {
		t.Errorf("ReadBackward case 1, got %v want %v", result1, expectedResult1)
	}
	if result2 != expectedResult2 {
		t.Errorf("ReadBackward case 2, got %v want %v", result2, expectedResult2)
	}
	if result3 != expectedResult3 {
		t.Errorf("ReadBackward case 3, got %v want %v", result3, expectedResult3)
	}
}

func TestRotorRoundTrip(t *testing.T) {
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE

	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ

	// Alpha: ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// Base1: EKMFLGDQVZNTOWYHXUSPAIBRCJ

	// Rotor Base1:EKMFLGDQVZNTOWYHXUSPAIBRCJ
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base2:AJDKSIRUXBLHWTMCQGZNPYFVOE
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	// Rotor Base3:BDFHJLCPRTXVZNYEIWGAKMUSQO
	// Alphabet:   ABCDEFGHIJKLMNOPQRSTUVWXYZ

	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring1, baseWiring2, baseWiring3)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}
	reflector := enigma.InitReflector(reflBaseWiring)

	forwardRead := enigma.ReadForward(rotors, "A")
	expForwardRead := "Z"

	if forwardRead != expForwardRead {
		t.Errorf("Forward read failed, got %v, want %v", forwardRead, expForwardRead)
	}

	reflectorRead := enigma.ReadReflector(reflector, forwardRead)
	expReflectorRead := "J"

	if reflectorRead != expReflectorRead {
		t.Errorf("Reflector read failed, got %v, want %v", reflectorRead, expReflectorRead)
	}

	backwardRead := enigma.ReadBackward(rotors, reflectorRead)
	expBackwardRead := "X"

	if backwardRead != expBackwardRead {
		t.Errorf("Backward read failed, got %v, want %v", backwardRead, expBackwardRead)
	}
}

// TestRotorRoundTripReflexive will encrypt and then decrypt
// a letter, which should result in the same letter output
func TestRotorRoundTripReflexive(t *testing.T) {
	// Initialize with same rotor settings as TestRotorRoundTrip
	baseWiring1 := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
	baseWiring2 := "AJDKSIRUXBLHWTMCQGZNPYFVOE"
	baseWiring3 := "BDFHJLCPRTXVZNYEIWGAKMUSQO"
	reflBaseWiring := "EKMFLGDQVZNTOWYHXUSPAIBRCJ"

	rotor1, rotor2, rotor3 := enigma.InitRotors3(baseWiring1, baseWiring2, baseWiring3)
	rotors := []*enigma.Rotor{rotor1, rotor2, rotor3}
	reflector := enigma.InitReflector(reflBaseWiring)

	forwardReadEncrypt := enigma.ReadForward(rotors, "A")
	reflectorReadEncrypt := enigma.ReadReflector(reflector, forwardReadEncrypt)
	backwardReadEncrypt := enigma.ReadBackward(rotors, reflectorReadEncrypt)
	expBackwardReadEncrypt := "X"

	if backwardReadEncrypt != expBackwardReadEncrypt {
		t.Errorf("Encrypt failed, got %v, want %v", backwardReadEncrypt, expBackwardReadEncrypt)
	}

	forwardReadDecrypt := enigma.ReadForward(rotors, "X")
	reflectorReadDecrypt := enigma.ReadReflector(reflector, forwardReadDecrypt)
	backwardReadDecrypt := enigma.ReadBackward(rotors, reflectorReadDecrypt)
	expBackwardReadDecrypt := "A"

	if backwardReadDecrypt != expBackwardReadDecrypt {
		t.Errorf("Decrypt failed, got %v, want %v", backwardReadDecrypt, expBackwardReadDecrypt)
	}
}
