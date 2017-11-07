package lz78

import (
	"fmt"
)

//Decompress ...
func Decompress(s string) (string, error) {

	// initialize empty output
	output := ""

	// integer index for dic map, must be integer for autoincrement
	mapIndex := 1

	dic := dictionary{}
	w := &dictEntry{}

	// lazy init for dictionary
	if dic.m == nil {
		dic.m = make(map[string]*dictEntry)
	}

	// iterate 7 by 7
	for i := 7; i <= len(s); i = i + 7 {

		idx := string(s[i-6 : i-2])
		k := string(s[i-2 : i-1])

		// create string index for dictionary map
		strMapIndex := fmt.Sprintf("%04d", mapIndex)

		// if idx == "0000" no search for value
		if idx != "0000" {
			dicPoint := dic.m[idx]
			w.s = dicPoint.s
		}

		dic.m[strMapIndex] = &dictEntry{w, k, w.s + k}

		// incremetn integer index for dic map
		mapIndex++

		// add to output
		output = output + w.s + k

		// create new instance of dictEntry
		w = &dictEntry{}
	}

	return output, nil
}
