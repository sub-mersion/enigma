// Pacakge enigma emulates a 3-rotor german Enigma encrypting machine.
package enigma

type (
	Plugboard [26]int
	Reflector [26]int

	Rotor struct {
		ForwardWiring  [26]int
		BackwardWiring [26]int
		NotchIndex     int
	}

	// Key is an enigma encrypting key, specifying the initial machine's settings.
	Key struct {
		// PluboardSetting is given as a string of letter pairs.
		// Example: "AR TY EP IS"
		PluboardSetting string
		Reflector       Reflector

		// Rotors is a 3-array of Rotors, orderer FROM RIGHT TO LEFT, namely the rotor at
		// index 0 is the most far right one on a real enigma machine.
		Rotors [3]Rotor

		// StartingPos and RingSettings are given as 3-arrays of capital letters.
		// Example: "[3]rune{'A', 'A', 'A'}"
		StartingPos  [3]rune
		RingSettings [3]rune
	}

	// Machine represents an enigma machine with a given key. Holding the
	// rings' positions during execution, it is the only stateful type.
	Machine struct {
		key Key
		// RotorPos holds the rotor positions. It is exported so that
		// they can be rotated prior to operation.
		RotorPos [3]int
	}
)
