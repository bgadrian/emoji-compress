package lz78

import "testing"

func TestDecompressBasic(t *testing.T) {
	in := "|0000a||0000b||0001b||0000c||0002a|"
	out := "ababcba"

	r, err := Decompress(in)
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}
