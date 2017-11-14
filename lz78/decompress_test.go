package lz78

import (
	"fmt"
	"log"
	"testing"
)

type casesInOut []inOut

type inOut struct {
	in  string
	out string
}

func TestDecompressInOut(t *testing.T) {

	c := casesInOut{
		{
			in:  "|0000a||0000b||0001b||0000c||0002a|",
			out: "ababcba",
		},
		{
			in:  "|0000a||0000b||0001b||0003c|",
			out: "abababc",
		},
		{
			in:  "|0000|||0000b||0001b||0003c|",
			out: "|b|b|bc",
		},
	}

	for _, e := range c {
		r, err := Decompress(e.in)
		if err != nil {
			t.Error(err)
		}

		if r != e.out {
			t.Errorf("Expected %s, got %s", e.out, r)
		}
	}
}

func ExampleDecompressString() {
	in := "|0000a||0000b||0001b||0000c||0002a|"
	out, err := Decompress(in)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("Decompressed string: %s", out)
	// Output: Decompressed string: ababcba
}
