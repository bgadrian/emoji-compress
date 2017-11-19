# LZ78 compressing algorithm [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/emoji-compress/lz78)
The package LZ78 can be used to compress and decompress strings into something nice: emojis.
 
 [LZ78](https://en.wikipedia.org/wiki/LZ77_and_LZ78)
  is a lossless data compression algorithm, which form the basis of several ubiquitous compression schemes, 
  including GIF and the DEFLATE algorithm used in PNG and ZIP.
  It has a simple [algorithm](http://www.stringology.org/DataCompression/lz78/index_en.html) which consist of finding repeating phrases (sequences of characters) and storing them in a tree like dictionary.
 
### Demo
We have built a full working demo at [emoji-compress.com](https://emoji-compress.com/) ‼

### How
* *Compress* function will iterate a string (slice of byte) character by character and  create a dictionary embedded in the output.
* *Decompress* function will use the embedded dictionary to recompose original string.

### Limitations:
* the emojis database has only around 1000 unique emojis, so the maximum length of a source text is limited
* the output length of the text could be bigger than the origina, if no sequences of characters repeat

## Example
```go
  // Compress
	in := "Play with emojis!"
	out, err := CompressString(in)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", out)
	// Output: 😀P😀l😀a😀y😀 😀w😀i😀t😀h😃e😀m😀o😀j😅s😀!

  // Decompress
	in := "😀P😀l😀a😀y😀 😀w😀i😀t😀h😃e😀m😀o😀j😅s😀!"
	out, err := DecompressString(in)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("%s", out)
	// Output: Play with emojis!
```
This package has unit tests, GoDoc and Examples.

### Resources:
* https://de.wikipedia.org/wiki/LZ78
* https://w3.ual.es/~vruiz/Docencia/Apuntes/Coding/Text/02-string_encoding/03-LZ78/index.html
* http://compressions.sourceforge.net/LempelZiv.html
* http://www.stringology.org/DataCompression/lz78/index_en.html
* https://unicode.org/emoji/charts/full-emoji-list.html

### About
This package is part of a [group of emoji-related encoding and compression algorithms](https://github.com/bgadrian/emoji-compress) built for fun and academic purposes in Go.

Copyright (c) 2017 [B.G.Adrian](https://coder.today) & @Davidescus

