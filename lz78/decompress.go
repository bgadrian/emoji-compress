package lz78

import (
	"io"
)

//Decompress ...
func Decompress(s string) (output string, err error) {

	dec := newDecoder(s)
	var phrase string
	var er error

	for err == nil {
		phrase, er = dec.decode()
		if er != nil {
			//EOF is a natural error, we all reach to an end
			if er != io.EOF {
				err = er
			}
			return
		}
		output = output + phrase
	}

	// // iterate 7 by 7
	// for i := 7; i <= len(s); i = i + 7 {
	// 	if string(s[i-7:i-6]) != "|" || string(s[i-1:i]) != "|" {
	// 		return "", fmt.Errorf("Malformed archive when decompressing, step %s is '%s'",
	// 			strconv.Itoa(i/7), s[i-7:i])
	// 	}
	// 	idx := string(s[i-6 : i-2])
	// 	k := string(s[i-2 : i-1])

	// 	// create string index for dictionary map
	// 	strMapIndex := fmt.Sprintf("%04d", mapIndex)

	// 	// if idx == "0000" no search for value
	// 	if idx != "0000" {
	// 		dicPoint := dic.m[idx]
	// 		if dicPoint == nil {
	// 			return "", fmt.Errorf("error when decoded, required idx %s but is missing.Dic size:%d",
	// 				idx, len(dic.m))
	// 		}
	// 		w.s = dicPoint.s
	// 	}

	// 	dic.m[strMapIndex] = &dictEntry{w, k, w.s + k}

	// 	// incremetn integer index for dic map
	// 	mapIndex++

	// 	// add to output
	// 	output = output + w.s + k

	// 	// create new instance of dictEntry
	// 	w = &dictEntry{}
	// }

	return output, nil
}
