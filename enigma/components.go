package enigma

var (
	//                   ABCDEFGHIJKLMNOPQRSTUVWXYZ
	RotorI   = newRotor("EKMFLGDQVZNTOWYHXUSPAIBRCJ", 'Q')
	RotorII  = newRotor("AJDKSIRUXBLHWTMCQGZNPYFVOE", 'E')
	RotorIII = newRotor("BDFHJLCPRTXVZNYEIWGAKMUSQO", 'V')
	RotorIV  = newRotor("ESOVPZJAYQUIRHXLNFTGKDCMWB", 'J')
	RotorV   = newRotor("VZBRGITYUPSDNHLXAWMJQOFECK", 'Z')

	//                         ABCDEFGHIJKLMNOPQRSTUVWXYZ
	ReflectorB = newReflector("YRUHQSLDPXNGOKMIEBFZCWVJAT")
	ReflectorC = newReflector("FVPJIAOYEDRZXWGCTKUQSBNMHL")
)
