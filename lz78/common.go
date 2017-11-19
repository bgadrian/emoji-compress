/*Package lz78
The package LZ78 can be used to compress and decompress strings into something nice: emojis.
The result of compressing with LZ78 will have more bytes in length than original string if it used for a small string, so not every time is an efficient compression method, is just a cooler and friendlier method that use emojis.*/
package lz78

import (
	"errors"
)

type dictEntry struct {
	w *dictEntry //a
	k string     //b
	s string     //ab
}

type dictionary struct {
	m map[string]*dictEntry
}

func (d *dictionary) contains(w *dictEntry, k string) (*dictEntry, bool) {
	wk := w.s + k
	wkAddress, ok := d.m[wk]
	return wkAddress, ok
}

func (d *dictionary) add(w *dictEntry, k string) error {
	if _, found := d.contains(w, k); found {
		return errors.New("duplicate entry, check Contains first")
	}

	if d.m == nil {
		d.m = make(map[string]*dictEntry)
	}

	wk := w.s + k
	d.m[wk] = &dictEntry{w, k, wk}

	return nil
}
