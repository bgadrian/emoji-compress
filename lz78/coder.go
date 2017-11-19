package lz78

import (
	"errors"
	"fmt"
	"io"

	"github.com/bgadrian/emoji-compressor/emojis"
)

//coder handles LZ78 result => output file and output file => LZ78
type coder struct {
	result   string
	tmp      map[string]string //"" => 😀 , "a" => 😬
	iterator emojis.Iterator
}

//output adds the new WK to the result
func (e *coder) output(ws, k string) error {
	//lazy init
	if e.tmp == nil {
		e.tmp = make(map[string]string, 200)
		zero, err := e.iterator.NextSingleRune()
		if err != nil {
			return err
		}
		e.tmp[""] = zero
	}

	uniqueAddress := e.tmp[""]
	ok := false
	if len(ws) > 0 {
		uniqueAddress, ok = e.tmp[ws]
		if ok == false {
			return errors.New(ws + " wasn't found in previous calls")
		}
	}

	nextEmoji, err := e.iterator.NextSingleRune()

	if err != nil {
		return err
	}

	e.tmp[ws+k] = nextEmoji

	e.result = e.result + fmt.Sprintf("%s%s", uniqueAddress, k)

	return nil
}

type decoder struct {
	archive  []rune
	tmp      map[string]*dictEntry //😀 => a
	iterator emojis.Iterator
}

// Create new instance of decoder and append it first value nil,
// becouse it need to start at 1
func newDecoder(archive []byte) *decoder {
	d := &decoder{}
	d.archive = []rune(string(archive))
	d.tmp = make(map[string]*dictEntry)
	d.iterator = emojis.Iterator{}

	zero, _ := d.iterator.NextSingleRune()
	d.tmp[zero] = nil
	return d
}

//decode from ...
func (e *decoder) decode() (phrase string, err error) {
	runesLeft := len(e.archive)
	if runesLeft == 0 {
		return "", io.EOF
	}

	if runesLeft < 2 {
		return "", fmt.Errorf("malformed archive, left: %s",
			string(e.archive))
	}

	current := e.archive[:2]
	e.archive = e.archive[2:]

	previousEmoji := string(current[0]) //the Emoji
	k := string(current[1])             //it's the A from "😬A"
	previousDictEntry, ok := e.tmp[previousEmoji]

	if ok == false {
		err = fmt.Errorf("Malformed archive, the address was not found yet. %s",
			previousEmoji)
		return
	}
	currentDictEntry := &dictEntry{k: k, s: k} //0000 w: nil, k: k, s:"" + k

	//prepend the value from the previous Address
	if previousDictEntry != nil {
		currentDictEntry.w = previousDictEntry
		currentDictEntry.s = currentDictEntry.w.s + currentDictEntry.k
	}

	//do not move this line before in the function
	currentEmoji, er := e.iterator.NextSingleRune()
	if er != nil {
		err = er
		return
	}

	e.tmp[currentEmoji] = currentDictEntry

	phrase = currentDictEntry.s
	return
}
