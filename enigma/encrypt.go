package enigma

import (
	"strings"
)

// Encrypt encrypts the input string starting from machine's current state. When
// called for the first time, it starts from the given key.
func (m *Machine) Encrypt(input string) string {

	runeInput := []rune(input)
	arr := make([]int, len([]rune(runeInput)))
	for i, r := range runeInput {
		arr[i] = int(r - 'A')
	}

	res := m.encryptFromInts(arr)

	var b strings.Builder
	for _, i := range res {
		b.WriteRune('A' + rune(i))
	}
	return b.String()
}

func (m *Machine) rotate() {
	if m.RotorPos[1] == m.key.Rotors[1].NotchIndex {
		m.RotorPos[2]++
		m.RotorPos[1]++
	}
	if m.RotorPos[0] == m.key.Rotors[0].NotchIndex {
		m.RotorPos[1]++
	}
	m.RotorPos[0]++
}

// encryptFromInts works with ints and perform the actual encryption.
func (m *Machine) encryptFromInts(input []int) []int {
	plugboard := newPlugboard(m.key.PluboardSetting)
	plugboardMapping := func(i int) int {
		return plugboard[i]
	}

	reflection := func(i int) int {
		return m.key.Reflector[i]
	}

	rotorsForward := func(i int) int {
		for j, r := range m.key.Rotors {
			shift := m.RotorPos[j] - int(m.key.RingSettings[j]-'A')
			i = mod26(r.ForwardWiring[mod26(i+shift)] - shift)
		}
		return i
	}

	rotorsBackward := func(i int) int {
		for j := 2; j >= 0; j-- {
			shift := m.RotorPos[j] - int(m.key.RingSettings[j]-'A')
			i = mod26(m.key.Rotors[j].BackwardWiring[mod26(i+shift)] - shift)
		}
		return i
	}

	// The whole encrypting function, except for rotors movements.
	// It actually performs 13 letter swaps on the alphabet.
	crypt := compose(
		plugboardMapping,
		rotorsForward,
		reflection,
		rotorsBackward,
		plugboardMapping,
	)

	res := make([]int, len(input))

	// Encryption loop
	for i, x := range input {
		if x < 0 || x >= 26 {
			res[i] = x
			continue
		}
		m.rotate()
		res[i] = crypt(x)
	}

	return res
}

// compose returns the mathemathical composition of the given functions.
// For example, compose(f, g, h)(i) is equivalent to h(g(f(i))).
func compose(fs ...func(int) int) func(int) int {
	return func(i int) int {
		for _, f := range fs {
			i = f(i)
		}
		return i
	}
}

// mod26 returns the mathematical modulo of i by 26. Result is in [0,26[.
func mod26(i int) int {
	return (i%26 + 26) % 26
}
