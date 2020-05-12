package enigma

import "math"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func offsetAdjust(rawBaseIndex int) (newIndex int) {
	// Mod denominator is the number of letters in the alphabet
	const alphaLetters float64 = 26

	if rawBaseIndex < 0 {
		newIndex = rawBaseIndex + 26
		return
	}

	newIndex = int(math.Mod(float64(rawBaseIndex), alphaLetters))
	return
}
