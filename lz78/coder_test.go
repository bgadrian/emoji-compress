package lz78

import (
	"io"
	"testing"
)

func TestCoderBasic(t *testing.T) {
	c := coder{}
	// a := dictEntry{nil, "a", "a"}
	// ab := dictEntry{&a, "b", "b"}

	err := c.output("", "a")
	if err != nil {
		t.Error(err)
	}

	result := "😀a" //0000a
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("", "b")
	if err != nil {
		t.Error(err)
	}

	result = "😀a😀b" //0000a000b
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("a", "b")
	if err != nil {
		t.Error(err)
	}

	result = "😀a😀b😬b" //ab[ab] 0000a0000b0001b
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("ab", "c")
	if err != nil {
		t.Error(err)
	}

	result = "😀a😀b😬b😂c" //abab[abc] 0000a 0000b 0001b 0003c
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}
}

func TestDecoderBasic(t *testing.T) {
	archive := "😀a😀b😬b😂c" //abab[abc] 0000a 0000b 0001b 0003c
	original := "abababc"
	d := newDecoder([]byte(archive))
	result := ""

	for true {
		phrase, err := d.decode()

		if err != nil {
			if err == io.EOF {
				break
			}
			t.Error(err)
			return
		}

		result = result + phrase
	}

	if result != original {
		t.Errorf("want %s got %s", original, result)
	}
}
