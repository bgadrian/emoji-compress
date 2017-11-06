package emojis

import (
	"errors"
	"unicode"
)

//Database An iterator over all the emoji database. Fetch 1 at a time.
type Database struct {
	c int
}

//ErrFinished Fetch returns this when the iterator is finished
const ErrFinished = "we ran out of emojis :("

//Fetch Return 1 emoji at a time from the "database".
func (e *Database) Fetch() (string, error) {
	if e.c >= len(inlinedb) {
		return "", errors.New(ErrFinished)
	}

	emoji := inlinedb[e.c]
	e.c++

	//TODO temporary fix until we find a lib/way to detect
	//composed emojis for decompression
	// 	https://en.wikipedia.org/wiki/Zero-width_joiner
	//  https://en.wikipedia.org/wiki/Variation_Selectors_(Unicode_block)
	if len([]rune(emoji)) > 1 {
		// log.Printf("skipping %d", []rune(emoji))
		return e.Fetch()
	}

	return emoji, nil
}

//IsEmoji Detect if a rune is reserved character.
func IsEmoji(r rune) bool {

	//TODO check only for the emojis from INLINEDB, not all Symbols
	//http://www.fileformat.info/info/unicode/category/So/index.htm
	return unicode.IsOneOf([]*unicode.RangeTable{
		unicode.So,
	}, r)

	// emotes := unicode.Ran
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
