package lz78

import (
	"testing"
)

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

func TestCompressRepetableString(t *testing.T) {
	in := "-!@-!@-!@-!@-!@-!@"
	out := "|0000-||0000!||0000@||0001!||0003-||0002@||0004@||0007-||0002@|"

	r, err := Compress([]byte(in))
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}

// func ExampleCompress() {
// 	in := "Play with emojis!"
// 	out, err := Compress([]byte(in))
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Printf("%s", out)
// 	// Output: |0000P||0000l||0000a||0000y||0000 ||0000w||0000i||0000t||0000h||0005e||0000m||0000o||0000j||0007s||0000!|
// }
