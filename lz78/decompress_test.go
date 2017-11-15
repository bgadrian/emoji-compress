package lz78

import (
	"fmt"
	"log"
	"testing"
)

type table struct {
	name string
	in   string
	out  string
}

func TestDecompress(t *testing.T) {

	tt := []table{
		{"Test1", "|0000a||0000b||0001b||0000c||0002a|", "ababcba"},
		{"Test2", "|0000a||0000b||0001b||0003c|", "abababc"},
		{"Test3", "|0000|||0000b||0001b||0003c|", "|b|b|bc"},
	}

	for _, e := range tt {
		r, err := Decompress(e.in)
		if err != nil {
			t.Error(err)
		}

		if r != e.out {
			t.Errorf("In %s expected %s, got %s", e.name, e.out, r)
		}
	}
}

func ExampleDecompress() {
	in := "|0000P||0000l||0000a||0000y||0000 ||0000w||0000i||0000t||0000h||0005e||0000m||0000o||0000j||0007s||0000!|"
	out, err := Decompress(in)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", out)
	// Output: Play with emojis!
}
