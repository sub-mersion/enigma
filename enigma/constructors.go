package enigma

import "strings"

func newPlugboard(def string) Plugboard {
	cables := strings.Fields(def)
	settings := make([][2]int, len(cables))
	for i, c := range cables {
		rs := []rune(c)
		settings[i] = [2]int{int(rs[0] - 'A'), int(rs[1] - 'A')}
	}
	return newPlugboardFromInts(settings)
}

func newPlugboardFromInts(settings [][2]int) Plugboard {
	var plugboard Plugboard
	for i := 0; i < 26; i++ {
		plugboard[i] = i
	}
	for _, wire := range settings {
		plugboard[wire[0]], plugboard[wire[1]] = plugboard[wire[1]], plugboard[wire[0]]
	}
	return plugboard
}

func newReflector(def string) Reflector {
	var arr Reflector
	for i, r := range def {
		arr[i] = int(r - 'A')
	}
	return arr
}

func newRotor(def string, notchIndex rune) Rotor {
	var fArr [26]int
	for i, r := range def {
		fArr[i] = int(r - 'A')
	}
	var bArr [26]int
	for i, x := range fArr {
		bArr[x] = i
	}
	return Rotor{
		ForwardWiring:  fArr,
		BackwardWiring: bArr,
		NotchIndex:     int(notchIndex - 'A'),
	}
}

// NewMachine returns an enigma Machine instance initialized with the given Key k.
func NewMachine(k Key) *Machine {
	var startingPos [3]int
	for i, p := range k.StartingPos {
		startingPos[i] = int(p - 'A')
	}

	return &Machine{
		key:      k,
		RotorPos: startingPos,
	}
}
