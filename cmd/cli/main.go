package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sub-mersion/enigma/enigma"
	"github.com/urfave/cli/v2"
)

var k = enigma.Key{
	Rotors:          [3]enigma.Rotor{enigma.RotorIII, enigma.RotorII, enigma.RotorI}, // From left to right
	Reflector:       enigma.ReflectorB,
	StartingPos:     [3]rune{'A', 'A', 'A'},
	RingSettings:    [3]rune{'A', 'A', 'A'},
	PluboardSetting: "",
}

func main() {
	app := cli.App{
		Name:  "enigma",
		Usage: "encrypt and decrypt text emulating a 3-rotor german enigma machine -- USE HARDOCDED KEY FOR NOW",
		Action: func(c *cli.Context) error {
			log.SetOutput(os.Stderr)
			m := enigma.NewMachine(k)
			input, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			output := m.Encrypt(string(input))
			fmt.Println(output)
			fmt.Println("input was", string(input))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
