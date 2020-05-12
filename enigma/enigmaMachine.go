package enigma

// Rotate function rotates a roter one notch
func (rotor *Rotor) Rotate() {
	rotor.offset++
}
