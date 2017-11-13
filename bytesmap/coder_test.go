package bytesmap

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func ExampleEncodeString() {
	ugly := []string{
		"127.0.0.1",
		"ZW1vamk=", //base64 for "emoji"
		"5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", //sha1 for "password"
	}

	for _, s := range ugly {
		emojified, err := EncodeString(s)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s => %s\n", s, emojified)
	}
	// Output: 127.0.0.1 => 🙇🙈🙍🙀🙆🙀🙆🙀🙇
	//ZW1vamk= => 🚇🚃🙇🚾🚕🚬🚪✉
	//5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8 => 🙋🚗🚕🚕🙌🙇🚢🙊🚙🙏🚗🙏🙉🚤🙉🚤🙆🙌🙎🙈🙈🙋🙆🚗🙌🚙🚤🙎🙉🙉🙇🚗🙍🚢🚢🙌🙎🚤🚚🙎
}

func ExampleDecodeString() {
	beauty := []string{
		"🙇🙈🙍🙀🙆🙀🙆🙀🙇", //127.0.0.1
		"🚇🚃🙇🚾🚕🚬🚪✉",  //base64 for "emoji"
		"🙋🚗🚕🚕🙌🙇🚢🙊🚙🙏🚗🙏🙉🚤🙉🚤🙆🙌🙎🙈🙈🙋🙆🚗🙌🚙🚤🙎🙉🙉🙇🚗🙍🚢🚢🙌🙎🚤🚚🙎", //sha1 for "password"
	}
	for _, s := range beauty {
		original, err := DecodeString(s)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s => %s\n", s, original)
	}
	// Output: 🙇🙈🙍🙀🙆🙀🙆🙀🙇 => 127.0.0.1
	//🚇🚃🙇🚾🚕🚬🚪✉ => ZW1vamk=
	//🙋🚗🚕🚕🙌🙇🚢🙊🚙🙏🚗🙏🙉🚤🙉🚤🙆🙌🙎🙈🙈🙋🙆🚗🙌🚙🚤🙎🙉🙉🙇🚗🙍🚢🚢🙌🙎🚤🚚🙎 => 5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8
}

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
		"A fost odată ca-n poveşti,\nA fost ca niciodată.\nDin rude mari împărăteşti,\nO prea frumoasă fată.",
		//from here https://golang.org/src/unicode/utf8/utf8_test.go
		"abcd",
		"☺☻☹",
		"日a本b語ç日ð本Ê語þ日¥本¼語i日©",
		"日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©",
		"\x80\x80\x80\x80",
	}

	for _, s := range table {
		err := compareInOut(s)
		if err != nil {
			t.Error(err)
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

//taken from the Unicode package tests

var upperTest = []rune{
	0x41,
	0xc0,
	0xd8,
	0x100,
	0x139,
	0x14a,
	0x178,
	0x181,
	0x376,
	0x3cf,
	0x13bd,
	0x1f2a,
	0x2102,
	0x2c00,
	0x2c10,
	0x2c20,
	0xa650,
	0xa722,
	0xff3a,
	0x10400,
	0x1d400,
	0x1d7ca,
}

var notupperTest = []rune{
	0x40,
	0x5b,
	0x61,
	0x185,
	0x1b0,
	0x377,
	0x387,
	0x2150,
	0xab7d,
	0xffff,
	0x10000,
}

var letterTest = []rune{
	0x41,
	0x61,
	0xaa,
	0xba,
	0xc8,
	0xdb,
	0xf9,
	0x2ec,
	0x535,
	0x620,
	0x6e6,
	0x93d,
	0xa15,
	0xb99,
	0xdc0,
	0xedd,
	0x1000,
	0x1200,
	0x1312,
	0x1401,
	0x2c00,
	0xa800,
	0xf900,
	0xfa30,
	0xffda,
	0xffdc,
	0x10000,
	0x10300,
	0x10400,
	0x20000,
	0x2f800,
	0x2fa1d,
}

var notletterTest = []rune{
	0x20,
	0x35,
	0x375,
	0x619,
	0x700,
	0x1885,
	0xfffe,
	0x1ffff,
	0x10ffff,
}

// Contains all the special cased Latin-1 chars.
var spaceTest = []rune{
	0x09,
	0x0a,
	0x0b,
	0x0c,
	0x0d,
	0x20,
	0x85,
	0xA0,
	0x2000,
	0x3000,
}
var testDigit = []rune{
	0x0030,
	0x0039,
	0x0661,
	0x06F1,
	0x07C9,
	0x0966,
	0x09EF,
	0x0A66,
	0x0AEF,
	0x0B66,
	0x0B6F,
	0x0BE6,
	0x0BEF,
	0x0C66,
	0x0CEF,
	0x0D66,
	0x0D6F,
	0x0E50,
	0x0E59,
	0x0ED0,
	0x0ED9,
	0x0F20,
	0x0F29,
	0x1040,
	0x1049,
	0x1090,
	0x1091,
	0x1099,
	0x17E0,
	0x17E9,
	0x1810,
	0x1819,
	0x1946,
	0x194F,
	0x19D0,
	0x19D9,
	0x1B50,
	0x1B59,
	0x1BB0,
	0x1BB9,
	0x1C40,
	0x1C49,
	0x1C50,
	0x1C59,
	0xA620,
	0xA629,
	0xA8D0,
	0xA8D9,
	0xA900,
	0xA909,
	0xAA50,
	0xAA59,
	0xFF10,
	0xFF19,
	0x104A1,
	0x1D7CE,
}

func TestUtf8Chars(t *testing.T) {
	var all [][]rune
	all = append(all, testDigit)
	all = append(all, spaceTest)
	all = append(all, notletterTest)
	all = append(all, letterTest)
	all = append(all, notupperTest)
	all = append(all, upperTest)

	for _, tests := range all {
		for _, r := range tests {
			err := compareInOut(string(r))
			if err != nil {
				t.Errorf("failed for utf8 %+q err: %v", r, err)
				break
			}
		}
	}
}
