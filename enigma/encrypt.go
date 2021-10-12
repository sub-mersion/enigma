package enigma

import (
	"log"
	"strings"
)

// Encrypt encrypts the input string starting from machine's current state. When
// called for the first time, it starts from the given key.
func (m *Machine) Encrypt(input string) string {
	arr := make([]int, len(input))
	for i, r := range input {
		arr[i] = int(r - 'A')
	}
	res := m.encryptFromInts(arr)

	var b strings.Builder
	for _, i := range res {
		b.WriteRune('A' + rune(i))
	}
	return b.String()
}

// encryptFromInts works with ints and perform the actual encryption.
func (m *Machine) encryptFromInts(input []int) []int {
	plugboard := newPlugboard(m.key.PluboardSetting)
	plugboardMapping := func(i int) int {
		log.Printf("Plugboard: %c\n", rune('A'+plugboard[i]))
		return plugboard[i]
	}

	reflection := func(i int) int {
		log.Printf("Reflector forward: %c\n", rune('A'+m.key.Reflector[i]))
		return m.key.Reflector[i]
	}

	rotorsForward := func(i int) int {
		for j, r := range m.key.Rotors {
			shift := m.RotorPos[j] - int(m.key.RingSettings[j]-'A')
			i = mod26(r.ForwardWiring[mod26(i+shift)] - shift)
			log.Printf("Rotors %d forward: %c\n", j, rune('A'+i))
		}
		return i
	}

	rotorsBackward := func(i int) int {
		for j := 2; j >= 0; j-- {
			shift := m.RotorPos[j] - int(m.key.RingSettings[j]-'A')
			i = mod26(m.key.Rotors[j].BackwardWiring[mod26(i+shift)] - shift)
			log.Printf("Rotors %d backward: %c\n", j, rune('A'+i))

		}
		return i
	}

	rotate := func() {
		if m.RotorPos[1] == m.key.Rotors[1].NotchIndex {
			m.RotorPos[2]++
			m.RotorPos[1]++
		}
		if m.RotorPos[0] == m.key.Rotors[0].NotchIndex {
			m.RotorPos[1]++
		}
		m.RotorPos[0]++
	}

	// The whole encrypting function, except for rotors movements
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
		log.Printf("Input is: %c", rune('A'+x))
		rotate()
		if x < 0 || x > 26 {
			continue
		}
		res[i] = crypt(x)
	}

	return res
}

// mod26 returns the mathematical modulo of i by 26. Result is in [0,26[.
func mod26(i int) int {
	return (i%26 + 26) % 26
}
