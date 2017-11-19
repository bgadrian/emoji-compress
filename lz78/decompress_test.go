package lz78

import (
	"fmt"
	"log"
	"testing"
)

func TestDecompress(t *testing.T) {

	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Decompress basic", "😀a😀b😬b😀c😁a", "ababcba"},
		{"Decompress more", "😀a😀b😬b😂c", "abababc"},
		{"Decompress repeat", "😀|😀b😬b😂c", "|b|b|bc"},
	}

	for _, e := range tt {
		r, err := DecompressString(e.in)
		if err != nil {
			t.Error(err)
		}

		if r != e.out {
			t.Errorf("In %s expected %s, got %s", e.name, e.out, r)
		}
	}
}

func ExampleDecompress() {
	in := "😀P😀l😀a😀y😀 😀w😀i😀t😀h😃e😀m😀o😀j😅s😀!"
	out, err := DecompressString(in)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", out)
	// Output: Play with emojis!
}
