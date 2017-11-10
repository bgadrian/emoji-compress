package lz78

import "testing"
import "strings"

import "github.com/sergi/go-diff/diffmatchpatch"
import "fmt"
import "io/ioutil"

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

func TestDecompressMultipleSequnces(t *testing.T) {
	in := "|0000a||0000b||0001b||0003c|"
	out := "abababc"

	r, err := Decompress(in)
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
	}
}

func TestDecompressPipeCharacter(t *testing.T) {
	in := "|0000|||0000b||0001b||0003c|"
	out := "|b|b|bc"

	r, err := Decompress(in)
	if err != nil {
		t.Error(err)
	}

	if r != out {
		t.Errorf("Expected %s, got %s", out, r)
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

func TestHeavy(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Recovered in TestHeavy", r)
		}
		return
	}()
	s, e := ioutil.ReadFile("big-poetry.txt")
	if e != nil {
		t.Error("Cant read file")
	}

	err := compareInOut(string(s))
	if err != nil {
		t.Error(err)
	}
}

func compareInOut(s string) error {
	c, err := CompressString(s)
	if err != nil {
		return err
	}

	d, err := Decompress(c)
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
