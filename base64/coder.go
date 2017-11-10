package base64

import (
	"bytes"
	"encoding/base64"
)

const encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

//TODO tweak the base64 package to worh with larger blocks
//4 byes each
// const encodeEmoji = "🤠🤗😏😶😐😑😒🙄🤔🤥😳😞😟😠😡😔😕🙁😣😖😫😩😤😮😱😨😰😯😦😧😢😥🤤😪😓😭😵😲🤐🤢🤧😷🤒🤕😴💤💩😈👿👹👺💀👻👽🤖😺😸😹😻😼😽🙀😿😾"
//Unicode 1 emojis - 3bytes each
// const encodeEmoji = "☀☁☂☃☄★☆☇☈☉☊☋☌☍☎☏☐☑☒☓☔☕☖☗☘☙☚☛☜☝☞☟☠☡☢☣☤☥☦☧☨☩☪☫☬☭☮☯♔♕♖♗♘♙♚♛♜♝♞♟♠♡♢♣" //♤♥♦♧♨♩♬♪♫

//EncodeEmoji 64bytes full of love and emojis to be used in base64
var EncodeEmoji *base64.Encoding

//EncodeString Encode a string to an emoji base64 version
func EncodeString(s string) ([]byte, error) {
	return Encode([]byte(s))
}

//Encode Encode string []byte  to an emoji base64 version
func Encode(input []byte) (output []byte, err error) {
	var b bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &b)
	_, err = encoder.Write(input)
	encoder.Close()
	output = b.Bytes()
	return
}

func init() {
	EncodeEmoji = base64.NewEncoding(encodeEmoji)
}
