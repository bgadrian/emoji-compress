package lz78

import (
	"io"
)

//Decompress ...
func Decompress(s string) (output string, err error) {
	dec := newDecoder(s)
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
