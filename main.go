package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
)

var m = map[string]int{"one": 1, "two": 2, "three": 3}

func main() {
	b := new(bytes.Buffer)

	e := gob.NewEncoder(b)

	// Encoding the map
	err := e.Encode(m)
	if err != nil {
		panic(err)
	}

	savePats("zzz", m)

	var decodedMap map[string]int
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	loadeddMap := loadPats("zzz")

	// Ta da! It is a map!
	fmt.Printf("%#v\n", decodedMap)
	fmt.Printf("%#v\n", loadeddMap)
}

func savePats(file string, pats map[string]int) {
	f, err := os.Create(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(pats); err != nil {
		panic("cant encode")
	}
}

func loadPats(file string) (pats map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err := enc.Decode(&pats); err != nil {
		panic("cant decode")
	}
	return pats
}
