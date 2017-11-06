package lz78

import (
	"fmt"
	"strings"
)

//Decompress ...
func Decompress(s string) (string, error) {

	// replace all ||
	s = strings.Replace(s, "|", "", -1)
	output := ""

	// integer index for dic map, must be integer for autoincrement
	mapIndex := 1

	dic := dictionary{}
	w := &dictEntry{}

	// lazy init for dictionary
	if dic.m == nil {
		dic.m = make(map[string]*dictEntry)
	}

	// iterate 5 by 5
	for i := 5; i <= len(s); i = i + 5 {

		idx := string(s[i-5 : i-1])
		k := string(s[i-1 : i])

		fmt.Println("Index: ", idx, " Character : ", k)

		// all 0000 idx must create new record in dictionary
		if idx == "0000" {

			// create string index for dictionary map
			strMapIndex := fmt.Sprintf("%04d", mapIndex)

			dic.m[strMapIndex] = &dictEntry{w, k, w.s + k}

			// incremetn integer index for dic map
			mapIndex++

			// create new instance of dictEntry
			w = &dictEntry{}

			// add to output
			output = output + k
			continue
		}

		// add key to buffer
		w.s = w.s + k

		pnt := dic.m[idx]
		output = output + pnt.s + k
	}

	return output, nil
}
