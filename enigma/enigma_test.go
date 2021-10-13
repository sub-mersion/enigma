package enigma

import (
	"fmt"
	"testing"
)

// TODO
// Test for text at theses pages
// https://www.codesandciphers.org.uk/enigma/enigma3.htm
// https://www.codesandciphers.org.uk/enigma/emachines/enigmad.htm
// https://fr.wikipedia.org/wiki/Enigma_(machine)

// Simple keys given as rotors starting position followed by ring settings.
var (
	AAA_AAA = Key{
		Rotors:          [3]Rotor{RotorIII, RotorII, RotorI},
		Reflector:       ReflectorB,
		StartingPos:     [3]rune{'A', 'A', 'A'},
		RingSettings:    [3]rune{'A', 'A', 'A'},
		PluboardSetting: "",
	}

	AAA_BBB = Key{
		Rotors:          [3]Rotor{RotorIII, RotorII, RotorI},
		Reflector:       ReflectorB,
		StartingPos:     [3]rune{'A', 'A', 'A'},
		RingSettings:    [3]rune{'B', 'B', 'B'},
		PluboardSetting: "",
	}

	JDK_DEA = Key{
		Rotors:          [3]Rotor{RotorIII, RotorII, RotorI},
		Reflector:       ReflectorB,
		StartingPos:     [3]rune{'J', 'D', 'K'},
		RingSettings:    [3]rune{'D', 'E', 'A'},
		PluboardSetting: "AZ EF GT IP",
	}
)

// TODO: improve the following example

// Encrypt the dummy text "AAAAA" with the rotors I, II and III from left to
// right with key AAA and ring settings AAA. Observe that the cypher is indeed
// symmetric.
func ExampleMachine_Encrypt() {
	m := NewMachine(AAA_AAA)
	fmt.Println(m.Encrypt("AAAAA"))
	m = NewMachine(AAA_AAA)
	fmt.Println(m.Encrypt("BDZGO"))
	// Output:
	// BDZGO
	// AAAAA
}

func TestCommonCypher(t *testing.T) {
	cases := []struct {
		key   Key
		input string
		want  string
	}{
		{AAA_AAA, "AAAAA", "BDZGO"},
		{AAA_AAA, "BDZGO", "AAAAA"},
		{AAA_BBB, "AAAAA", "EWTYX"},
		{AAA_BBB, "EWTYX", "AAAAA"},
	}

	for _, tt := range cases {
		m := NewMachine(tt.key)
		got := m.Encrypt(tt.input)
		if got != tt.want {
			t.Fatalf("got %q want %q", got, tt.want)
		}
	}
}

func TestSymmetry(t *testing.T) {
	input := "IMETATRAVELLERFROMANANTIQUELANDWHOSAIDTWOVASTANDTRUNKLESSLEGSOFSTONESTANDINTHEDESERTNEARTHEMONTHESANDHALFSUNKASHATTEREDVISAGELIESWHOSEFROWNANDWRINKLEDLIPANDSNEEROFCOLDCOMMANDTELLTHATITSSCULPTORWELLTHOSEPASSIONSREADWHICHYETSURVIVESTAMPEDONTHESELIFELESSTHINGSTHEHANDTHATMOCKEDTHEMANDTHEHEARTTHATFEDANDONTHEPEDESTALTHESEWORDSAPPEARMYNAMEISOZYMANDIASKINGOFKINGSLOOKONMYWORKSYEMIGHTYANDDESPAIRNOTHINGBESIDEREMAINSROUNDTHEDECAYOFTHATCOLOSSALWRECKBOUNDLESSANDBARETHELONEANDLEVELSANDSSTRETCHFARAWAY"
	m := NewMachine(JDK_DEA)
	cypher := m.Encrypt(input)
	m = NewMachine(JDK_DEA)
	got := m.Encrypt(cypher)
	if got != input {
		t.Fatalf("got\n%s\nwant\n%s\n", got, input)
	}
}

// TestNonUppercaseRune check that runes outside the range [A-Z] are streamed
// untouched to the output, as promised by a TODO in the README. It implies to
// rewrite some code in the Machine.Encrypt() method, as it does not work with
// rune in a satisfactory manner. Also, there should be a warning issue and we
// should check this as well.
func TestNonUppercaseRune(t *testing.T) {
	input := "abcd (§12-_*$^ùâĒ?~	}\t\n"
	m := NewMachine(JDK_DEA)
	got := m.Encrypt(input)
	if got != input {
		t.Fatalf("got %q want %q", got, input)
	}
}
