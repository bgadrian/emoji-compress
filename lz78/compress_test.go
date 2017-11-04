package lz78

import "testing"

func TestCompressBasic(t *testing.T) {
	in := "ababcba"
	out := "|0000a||0000b||0001b||0000c||0002a|"

	r, err := Compress([]byte(in))
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}

func TestCompressMore(t *testing.T) {
	in := "ababcbababaaaaaaaaa"
	out := "|0000a||0000b||0001b||0000c||0002a||0005b||0001a||0007a||0008a|"

	r, err := Compress([]byte(in))
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}
