package lz78

import (
	"errors"
	"fmt"
)

//coder handles LZ78 result => output file and output file => LZ78
type coder struct {
	result string
	tmp    map[string]int
}

//output adds the new WK to the result
func (e *coder) output(ws, k string) error {

	//lazy init
	if e.tmp == nil {
		e.tmp = make(map[string]int, 200)
	}

	//TODO replace with a real algorithm
	//The dictionary values can be encoded by one of entropy coders. It can be Huffman, Arithmetic, Universal Integer Coding etc.
	uniqueAddress := 0
	ok := false
	if len(ws) > 0 {
		uniqueAddress, ok = e.tmp[ws]
		if ok == false {
			return errors.New(ws + " wasn't found in previous calls")
		}
	}

	e.tmp[ws+k] = len(e.tmp) + 1
	e.result = e.result + fmt.Sprintf("|%04d%s|", uniqueAddress, k)

	return nil
}

//decode from ...
func (e *coder) decode() (dictionary, error) {

	return dictionary{}, nil
}
