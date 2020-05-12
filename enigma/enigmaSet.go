package enigma

import "fmt"

// SetRotorOffset sets the offset for the specified rotor
func (rotor *Rotor) SetRotorOffset(newOffset int) {

	if newOffset > 25 {
		fmt.Printf("The offset you have requested, %v, is higher than maximum of 25. It will remain at %v\n", newOffset, rotor.offset)
		return
	}

	rotor.offset = newOffset
}
