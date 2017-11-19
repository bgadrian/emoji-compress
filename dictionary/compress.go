package dictionary

import (
	"errors"
	"unicode"
	"unicode/utf8"

	"github.com/bgadrian/emoji-compress/emojis"
)

//CompressString is an overload for Compress.
func CompressString(s string) (*Result, error) {
	return Compress([]byte(s))
}

//Compress a string by replacing it's words with emojis.
//The result will contain a map of words/emojis that can be used to revert the process.
func Compress(original []byte) (r *Result, err error) {
	r = &Result{}
	r.Source = string(original)
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
	for utf8.RuneCount(original) > 0 {
		nextRune, size := utf8.DecodeRune(original)
		original = original[size:]
		isWordDelimiter := unicode.IsSpace(nextRune) || unicode.IsPunct(nextRune)
		isEmoji := emojis.IsEmoji(string(nextRune))

		if isEmoji {
			//it is a simple algorithm, you are asking for too much
			//TODO find a way to allow usage of emoji in the original text
			err = errors.New("Illegal character found: " + string(nextRune))
			return
		}

		if isWordDelimiter == false {
			word += string(nextRune)
			continue
		}
		//else is a word delimiter

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
