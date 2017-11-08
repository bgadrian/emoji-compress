package lz78

import (
	"unicode/utf8"
)

//CompressString ...
func CompressString(s string) (string, error) {
	return Compress([]byte(s))
}

//Compress ...
func Compress(input []byte) (string, error) {

	dic := dictionary{}
	coder := coder{}

	w := &dictEntry{}
	k := "" //symbol == rune == 1 character string

	for utf8.RuneCount(input) > 0 {
		rune, size := utf8.DecodeRune(input)
		input = input[size:]
		k = string(rune)

		wk, found := dic.contains(w, k)
		if found {
			w = wk
			continue
		}

		coder.output(w.s, k)
		dic.add(w, k)
		w = &dictEntry{} //nil
	}

	return coder.result, nil
}
