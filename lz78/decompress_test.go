package lz78

import "testing"

type casesInOut []inOut

type inOut struct {
	in  string
	out string
}

func TestDecompresInOut(t *testing.T) {

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
