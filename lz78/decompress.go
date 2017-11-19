package lz78

import (
	"io"
)

// DecompressString is an overload for Decompress.
func DecompressString(s string) (string, error) {
	return Decompress([]byte(s))
}

// Decompress will iterate input that has embeded dictionary inside of it
// and will recompose the original string
func Decompress(input []byte) (output string, err error) {
	dec := newDecoder(input)
	var phrase string
	var e error

	for true {
		phrase, e = dec.decode()
		if e != nil {
			//EOF is a natural error, we all reach to an end
			if e == io.EOF {
				return
			}
			err = e
			return
		}
		output = output + phrase
	}

	return
}
