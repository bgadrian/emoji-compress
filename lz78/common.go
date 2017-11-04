//Package lz78 TODO
//
//https://w3.ual.es/~vruiz/Docencia/Apuntes/Coding/Text/02-string_encoding/03-LZ78/index.html
//http://compressions.sourceforge.net/LempelZiv.html
//http://www.stringology.org/DataCompression/lz78/index_en.html
//https://unicode.org/emoji/charts/full-emoji-list.html
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
