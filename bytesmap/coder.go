package bytesmap

import "fmt"

//EncodeString overload of Encode, work with strings directly
func EncodeString(source string) (string, error) {
	bytes, err := Encode([]byte(source))
	return string(bytes), err
}

//Encode a stream of bytes (usually a string) into an emoji form.
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

//DecodeString overload for Decode
func DecodeString(encoded string) (string, error) {
	bytes, err := Decode([]byte(encoded))
	return string(bytes), err
}

//Decode an emoji encoded to it's previous form.
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
