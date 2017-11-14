# Bytesmap emoji encoding  [![GoDoc](https://godoc.org/golang.org/x/tools/cmd/godoc?status.svg)](https://godoc.org/github.com/bgadrian/emoji-compressor/bytesmap)
The package bytesmap can be used to encode strings into emojis.
Can be useful to "humanize" hard to remember/recognize texts such as
Hashes, keys, other encodes like base64, ip's and so on.

Depending on the text the result may have fewer characters,
but definately will have more bytes in length, so is not an
efficient compresison method, is just an encoding like base64,
but way cooler and friendlier.

### How
The algorithm is very simple: *it split the string in a series of bytes,
and map each byte by its value to an unique emoji.*
A byte can have only **255** possible values, so we only need
255 different **emojis** to encode ... anything.

### Source
The package is a Go port of @ayende 's [emoji encoder](https://ayende.com/blog/177729/emoji-encoding-a-new-style-for-binary-encoding-for-the-web). He also uses it to encode the Licenses for his product when they are sent to the customers.
>There are other advantages. This data is actually a 256 bits key for use in encryption. And you can actually show it to a user and have a reasonably good chance that they will be able to tell it apart from something else. It rely on the ability of humans to recognize shapes, but it will be very hard for them to actually tell someone your key.


## Example
```go
	ugly := []string{
		"127.0.0.1",
		"ZW1vamk=", //base64 for "emoji"
		"5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8", //sha1 for "password"
	}

	for _, s := range ugly {
		emojified, err := EncodeString(s)
		if err != nil {
			log.Panic(err)
		}
		fmt.Printf("%s => %s\n", s, emojified)
	}
	// Output: 127.0.0.1 => 🙇🙈🙍🙀🙆🙀🙆🙀🙇
	//ZW1vamk= => 🚇🚃🙇🚾🚕🚬🚪✉
	//5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8 => 🙋🚗🚕🚕🙌🙇🚢🙊🚙🙏🚗🙏🙉🚤🙉🚤🙆🙌🙎🙈🙈🙋🙆🚗🙌🚙🚤🙎🙉🙉🙇🚗🙍🚢🚢🙌🙎🚤🚚🙎
```
This package has unit tests, GoDoc and Examples.

### About
This package is part of a [group of emoji-related encoding and compression algorithms](https://github.com/bgadrian/emoji-compressor) built for fun and academic purposes in Go.

Copyright (c) 2017 [B.G.Adrian](https://coder.today) & @Davidescus
