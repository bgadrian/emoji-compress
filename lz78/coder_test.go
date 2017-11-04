package lz78

import "testing"

func TestCoderBasic(t *testing.T) {
	c := coder{}
	// a := dictEntry{nil, "a", "a"}
	// ab := dictEntry{&a, "b", "b"}

	err := c.output("", "a")
	if err != nil {
		t.Error(err)
	}

	result := "|0000a|"
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("", "b")
	if err != nil {
		t.Error(err)
	}

	result = "|0000a||0000b|"
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("a", "b")
	if err != nil {
		t.Error(err)
	}

	result = "|0000a||0000b||0001b|" //ab[ab]
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}

	err = c.output("ab", "c")
	if err != nil {
		t.Error(err)
	}

	result = "|0000a||0000b||0001b||0003c|" //abab[abc]
	if c.result != result {
		t.Errorf("Expected %s, got %s", result, c.result)
	}
}
