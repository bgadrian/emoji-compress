package dictionary

import (
	"errors"
	"unicode"
	"unicode/utf8"

	"github.com/bgadrian/emoji-compressor/emojis"
)

//Result Contains the original phrase, dictionary and other data for Compress and Decompress
type Result struct {
	Words   map[string]string `json:"dict"`
	Source  string            `json:"source"`
	Archive string            `json:"archive"`
	Ratio   float32           `json:"ratio"`
}

const min int = 2

//CompressString Compress a string using emojis instead of words.
func CompressString(s string) (*Result, error) {
	return Compress([]byte(s))
}

//Compress Compress a string using emojis instead of words.
func Compress(input []byte) (r *Result, err error) {
	r = &Result{}
	r.Source = string(input)
	r.Words = make(map[string]string, 100)
	r.Archive = ""

	word := ""
	db := emojis.Iterator{}

	step := func() {
		if len(word) <= min {
			return
		}
		//we have a word to replace
		emoji, ok := r.Words[word]
		if ok == false {
			emoji, err = db.NextSingleRune()
			if err != nil {
				return
			}
			r.Words[word] = emoji
		}
		word = ""
		r.Archive += string(emoji)
	}

	//go trough each character
	for utf8.RuneCount(input) > 0 {
		nextRune, size := utf8.DecodeRune(input)
		input = input[size:]
		isCompressable := unicode.IsLetter(nextRune) || unicode.IsDigit(nextRune)
		isEmoji := emojis.IsEmoji(string(nextRune))

		if isEmoji {
			//it is a simple algorithm, you are asking for too much
			//TODO find a way to allow usage of emoji in the original text
			err = errors.New("Illegal character found: " + string(nextRune))
			return
		}

		if isCompressable {
			word += string(nextRune)
			continue
		}
		//else is a non alfa numeric character

		//we have leftovers
		step()

		//we keep it as it is
		r.Archive += word + string(nextRune)
		word = ""
	}
	step()

	r.Ratio = float32(len(r.Archive)) / float32(len(r.Source))
	return
}
