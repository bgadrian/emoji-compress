/*Package bytesmap can be used to encode strings into emojis.
Can be useful to "humanize" hard to remember/recognize texts such as
Hashes, keys, other encodes like base64, ip's and so on.

Depending on the text the result may have fewer characters,
but definately will have more bytes in length, so is not an
efficient compresison method, is just an encoding like base64,
but way cooler and friendlier.

The algorithm is very simple: it split the string in a series of bytes,
and map each byte by its value to an unique emoji.
A byte can have only 255 possible values, so we only need
255 different emojis to encode ... anything. */
package bytesmap

import "fmt"

//EncodeString is an overload of Encode(), most of the cases you'll need this
//to avoid code duplication ([]byte <-> string).
func EncodeString(source string) (string, error) {
	bytes, err := Encode([]byte(source))
	return string(bytes), err
}

//Encode a stream of bytes (we presume it is a string) into an emoji form.
//Each possible byte value (0-255) is mapped to an unique emoji.
func Encode(source []byte) (encoded []byte, err error) {
	var r rune
	var asByte []byte

	max := len(db) - 1
	for _, b := range source {
		if b < 0 || int(b) > max {
			err = fmt.Errorf("found a byte outside of db %v", b)
			return
		}
		r = db[b]
		asByte = []byte(string(r))
		encoded = append(encoded, asByte...)
	}
	return
}

//DecodeString overload the Decode() function.
func DecodeString(encoded string) (string, error) {
	bytes, err := Decode([]byte(encoded))
	return string(bytes), err
}

//Decode an emoji encoded to it's original form.
func Decode(encoded []byte) (source []byte, err error) {
	var ok bool
	var b byte
	asString := string(encoded)

	for _, r := range asString {
		b, ok = reversedb[r]

		if ok == false {
			err = fmt.Errorf("found a missing byte from reversed db %v", r)
			return
		}

		source = append(source, b)
	}
	return
}
