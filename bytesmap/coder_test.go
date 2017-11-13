package bytesmap

import (
	"fmt"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestMapSanity(t *testing.T) {
	l := len(db)
	if l != 256 {
		t.Errorf("the map db must have 256 entries, it has %d",
			l)
	}

	//the index of C should be the same as value in reversed map
	for i, r := range db {
		ri := reversedb[r]

		if int(ri) != i {
			t.Errorf("revered db is malformed, exp %d got %d for %v",
				i, ri, r)
		}
	}
}

func TestEncodeBasic(t *testing.T) {
	for b := byte(0); b <= 255; b++ {
		o, err := Encode([]byte{b})

		if err != nil {
			t.Error(err)
			break
		}

		asString := string(o)
		asRunes := []rune(asString)
		if len(asRunes) != 1 {
			t.Errorf("1 byte should return 1 rune, got %v : %s for byte %d",
				asRunes, asString, b)
			break
		}

		emoji := db[int(b)]
		if asRunes[0] != emoji {
			t.Errorf("wrong emoji for byte %d: %v",
				b, asString)
			break
		}

		if b == 255 {
			break
		}
	}
}

func TestFullTableASCII(t *testing.T) {
	table := []string{
		"",
		" ",
		"123456789012",
		"~!@~!@~!@~!@~!@~!@",
		"Broasca are sau nu are mere?",
		//utf8 diacritics
		"A fost odată ca-n poveşti,\nA fost ca niciodată.\nDin rude mari împărăteşti,\nO prea frumoasă fată.",
		//TODO add more here
	}

	for _, s := range table {
		err := compareInOut(s)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFullDynamicASCII(t *testing.T) {
	from := 0
	to := 128
	c := []string{
		"a", "aa", "a?", "aa?",
		"aba", "abba", "bbaa",
		//TODO add more here
	}

Exit:
	for _, chars := range c {
		for i := from; i < to; i++ {
			source := strings.Repeat(chars, i)
			err := compareInOut(source)

			if err != nil {
				t.Error(err)
				break Exit
			}
		}
	}
}

func compareInOut(s string) error {
	c, err := EncodeString(s)
	if err != nil {
		return err
	}

	d, err := DecodeString(c)
	if err != nil {
		return err
	}

	if strings.Compare(s, d) != 0 {
		//because the texts are "aaabbababaaaba" it's very hard to see
		//the diff between source and output, so we're using a diff colored result
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(s, d, false)
		return fmt.Errorf("source malformed after decompress \nexp: %s \ndiff: %s \ngot: %s",
			s, dmp.DiffPrettyText(diffs), d)
	}

	return nil
}
