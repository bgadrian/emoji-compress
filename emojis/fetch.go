//Package emojis can be used as an iterator over most of the emoji in
//Unicode in a predefined (but random) order.
//The emojis are hardcoded for now, until we can find a better solution.
//Usage:
// db := Iterator{}
// db.Next() //returns "😀"
package emojis

import (
	"errors"
)

//Iterator An iterator over all the emoji database. Fetch 1 at a time.
type Iterator struct {
	pos int
}

//EOF Fetch functions use this error to gracefully alert that they
//reached the end.
var EOF = errors.New("EOF")

//Next server the next unique emoji from the internal DB.
func (e *Iterator) Next() (string, error) {
	if e.Done() {
		return "", EOF
	}

	emoji := inlinedb[e.pos]
	e.pos++

	return emoji, nil
}

//NextSingleRune Return the next emoji which consists of just 1 rune.
//Attention! All emojis runes have >= 2 bytes, 1 rune (character) doesn't mean 1 byte.
func (e *Iterator) NextSingleRune() (string, error) {
	emoji, err := e.Next()

	for len([]rune(emoji)) > 1 && err == nil {
		emoji, err = e.Next()
	}
	//TODO temporary fix until we find a lib/way to detect
	//composed emojis for decompression
	// 	https://en.wikipedia.org/wiki/Zero-width_joiner
	//  https://en.wikipedia.org/wiki/Variation_Selectors_(Unicode_block)

	return emoji, err
}

//Done Check if the iterator finished going trough all the emojis.
func (e *Iterator) Done() bool {
	return e.pos >= len(inlinedb)
}

//IsEmoji Detect if a rune is part of our internal DB.
func IsEmoji(r string) bool {
	_, ok := inlinemap[r]
	return ok
	//TODO check only for the emojis from INLINEDB, not all Symbols
	//http://www.fileformat.info/info/unicode/category/So/index.htm
	// switch value {
	//     case 0x1F600...0x1F64F, // Emoticons
	//         0x1F300...0x1F5FF, // Misc Symbols and Pictographs
	//         0x1F680...0x1F6FF, // Transport and Map
	//         0x2600...0x26FF,   // Misc symbols
	//         0x2700...0x27BF,   // Dingbats
	//         0xFE00...0xFE0F,   // Variation Selectors
	//         0x1F900...0x1F9FF,  // Supplemental Symbols and Pictographs
	//         65024...65039, // Variation selector
	//         8400...8447: // Combining Diacritical Marks for Symbols
	//         return true
}

func init() {
	inlinemap = make(map[string]struct{}, len(inlinedb))
	for _, e := range inlinedb {
		inlinemap[e] = struct{}{}
	}
}
