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

	saveFile("file.bin", m)

	var decodedMap map[string]int
	d := gob.NewDecoder(b)

	// Decoding the serialized data
	err = d.Decode(&decodedMap)
	if err != nil {
		panic(err)
	}

	loadeddMap := loadFile("file.bin")

	// Ta da! It is a map!
	fmt.Printf("%#v\n", decodedMap)
	fmt.Printf("%#v\n", loadeddMap)
}

func saveFile(file string, m map[string]int) {
	f, err := os.Create(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewEncoder(f)
	if err := enc.Encode(m); err != nil {
		panic("cant encode")
	}
}

func loadFile(file string) (m map[string]int) {
	f, err := os.Open(file)
	if err != nil {
		panic("cant open file")
	}
	defer f.Close()

	enc := gob.NewDecoder(f)
	if err := enc.Decode(&m); err != nil {
		panic("cant decode")
	}
	return m
}
