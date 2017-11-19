package lz78

import (
	"fmt"
	"log"
	"testing"
)

func TestCompress(t *testing.T) {

	tt := []struct {
		name string
		in   string
		out  string
	}{
		{"Compress basic", "abababc", "😀a😀b😬b😂c"},
		// {"Compress more", "ababcbababaaaaaaaaa", "😀a😀b0001b😀c0002a0005b0001a0007a0008a"},
		// {"Compress repeat", "-!@-!@-!@-!@-!@-!@", "😀-😀!😀@0001!0003-0002@0004@0007-0002@"},

		{"Single letter", "a", "😀a"},
		//{"Compress more", "ababcbababaaaaaaaaa", "😀😬😀b0001b😀c0002a0005b0001a0007a0008a"},
		//{"Compress repeat", "-!@-!@-!@-!@-!@-!@", "😀-😀!😀@0001!0003-0002@0004@0007-0002@"},
	}

	for _, e := range tt {
		r, err := Compress([]byte(e.in))
		if err != nil {
			t.Error(err)
		}

		if r != e.out {
			t.Errorf("%s expected %s, got %s", e.name, e.out, r)
		}
	}
}

func ExampleCompress() {
	in := "Play with emojis!"
	out, err := Compress([]byte(in))
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", out)
	// Output: 😀P😀l😀a😀y😀 😀w😀i😀t😀h😃e😀m😀o😀j😅s😀!
}
