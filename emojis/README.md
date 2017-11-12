# Emoji Iterator - Go package  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/emoji-compressor/emojis)
The package emojis can be used as an iterator over most of the emoji in Unicode in a predefined (but random) order.

Current count: **1183 emojis** from which **1011** have 1 rune length.

```go
db := Iterator{}
db.Next() //returns "😀"
db.Next() //returns "😬"
//repeat for 1181 times
db.Next() //returns "" and emojis.EOF error
```

The function **Iterator.NextSingleRune()** is used to fetch the next single-rune length emoji. The result is:
```
1 string (ex: "😀")
1 rune the string has len([]rune("😀")) == 1
>1 bytes. Even though is 1 character it has more bytes (see Unicode).
```

Emojis package has unit tests, GoDoc and Examples.
We use this package in our Encoding and Compressiong algorithms we have built for academic (and fun) purposes.

*The emojis are hardcoded for now, until we can find a better solution.*
