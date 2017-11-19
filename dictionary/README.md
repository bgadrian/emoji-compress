# Dictionary emoji encoding  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/emoji-compressor/dictionary)
Package dictionary is a small package that allows encoding (or compression) of strings by replacing each unique word with an emoji.

Each compress generates a new dictionary and an encoded version of the text (archive), based on the words found in the text.
If the original text doesn't have many repeating words, the "archive" will be longer than the original string.

The dictionary should be sent to the user/client so he can decode the string.

### Demo
We have built a full working demo at [emoji-compress.com](https://emoji-compress.com/) ‼

### Limitations:
* you cannot have emojis in the original text
* only works with a max of 1000 unique words
* (for now) compress generates a new dictionary for each text
* you have to use the same dictionary resulted from the Compress into the Decompress

### TODO:
* the ability to use a custom dictionary when Compressing

### How
The algorithm is very simple: tries to extract each word from the original text and replace it with an emoji.
A dictionary/map is generated along with the "Archive", to remember which word was replaced with each emoji.

The decompress process requires the "Archived" (encoded) version of the text, and the dictionary, in order to reverse the process.


## Example
```go
	//snippet of Sonnet 40 Take all my loves, my love, yea, take them all BY WILLIAM SHAKESPEARE
	sonnet := "Take all my loves, my love, yea, take them all:" +
		"\nWhat hast thou then more than thou hadst before?" +
		"\nNo love, my love, that thou mayst true love call—" +
		"\nAll mine was thine before thou hadst this more."

	result, err := CompressString(sonnet)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Archive: %s", result.Archive)
	j, err := json.Marshal(result.Words)
	fmt.Printf("\nDictionary: %s", j)

	// Output: Archive: 😀 😬 my 😁, my 😂, 🤣, 😃 😄 😬:
	// 😅 😆 😇 😉 😊 🙂 😇 🙃 😋?
	// No 😂, my 😂, 😌 😇 😍 😘 😂 😗—
	// 😙 😚 😜 😝 😋 😇 🙃 😛 😊.
	// Dictionary: {"All":"😙","Take":"😀","What":"😅","all":"😬","before":"😋","call":"😗","hadst":"🙃","hast":"😆","love":"😂","loves":"😁","mayst":"😍","mine":"😚","more":"😊","take":"😃","than":"🙂","that":"😌","them":"😄","then":"😉","thine":"😝","this":"😛","thou":"😇","true":"😘","was":"😜","yea":"🤣"}

```
This package has unit tests, GoDoc and Examples.

### About
This package is part of a [group of emoji-related encoding and compression algorithms](https://github.com/bgadrian/emoji-compressor) built for fun and academic purposes in Go.

Copyright (c) 2017 [B.G.Adrian](https://coder.today) & @Davidescus
