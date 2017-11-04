package lz78

import "testing"

func TestCompressBasic(t *testing.T) {
	in := "ababcba"
	out := "|0000a||0000b||0001b||0000c||0002a|"

	r, err := Compress(in)
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}
