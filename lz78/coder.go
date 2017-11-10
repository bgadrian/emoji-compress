package lz78

import (
	"errors"
	"fmt"
	"io"
	"strconv"
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

type decoder struct {
	archive []rune
	// dic     dictionary
	tmp []*dictEntry
}

func newDecoder(archive string) *decoder {
	d := &decoder{}
	d.archive = []rune(archive)
	// d.dic = make(map[string]*dictEntry, 10)
	d.tmp = append(d.tmp, &dictEntry{nil, "", ""})
	return d
}

//decode from ...
func (e *decoder) decode() (phrase string, err error) {
	runesLeft := len(e.archive)
	if runesLeft == 0 {
		return "", io.EOF
	}

	if runesLeft < 7 {
		return "", fmt.Errorf("Malformed archive, left: %s",
			string(e.archive))
	}

	current := e.archive[:7]
	e.archive = e.archive[7:]

	addressStr := string(current[1:5])
	k := string(current[5])
	address, er := strconv.Atoi(addressStr)

	if er != nil {
		err = er
		return
	}
	w := &dictEntry{k: k, s: k}

	//prepend the value from the previous Address
	if address > 0 {
		w.w = e.tmp[address]

		if w.w == nil {
			err = fmt.Errorf("malformed, address not yet found %s %s",
				addressStr, string(current))
			return
		}
		w.s = w.w.s + w.k
	}
	e.tmp = append(e.tmp, w)

	phrase = w.s
	return
}
