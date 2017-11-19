# Emoji-compress

**Emoji-compress** is an open-source library, written in Go, as a side-project. We have ported a few known basic encoding and compression algorithms and added the emoji twist. Some of the methods may result in smaller texts (as in the number of characters), but larger in bytes.

All packages have unit tests and GoDocumentation. The project also contains a demo web server and website, which aren't supposed to be uses in a (serious, non-emojified) production env.

## Demo
#### We have built a full working demo at [emoji-compress.com](https://emoji-compress.com/) ‼

## Packages
### [LZ78 compressing algorithm](../lz78/README.md)
LZ78 is a lossless data compression algorithm, which form the basis of several ubiquitous compression schemes,  including GIF and the DEFLATE algorithm used in PNG and ZIP.

```go
source := "“No heart is so hard as the timid heart.”"
archive := "😀“😀N😀o😀 😀h😀e😀a😀r😀t🤣i😀s🤣s😂 😃a😆d🤣a😊 😇h😄 😇i😀m😀i😀d🤣h😄a😆t😀.😀”"
```
See more [details here](../lz78/README.md).

### [Bytes map - encoding](../bytesmap/README.md)
It is a simple encoding method, it is use to "humanize" hard to remember/recognize texts such as Hashes, keys, other encodes like base64, ip's and so on.

```go
source := "127.0.0.1"
archive := "🙇🙈🙍🙀🙆🙀🙆🙀🙇"
```
See more [details here](../bytesmap/README.md).

### [Dictionary - encoding](../dictionary/README.md)
Package dictionary is a small package that allows encoding (or compression) of strings by replacing each unique word with an emoji. Each compress generates a new dictionary and an encoded version of the text (archive), based on the words found in the text.

```go
source := "“I felt happy because I saw the others were happy and because I knew I should feel happy, but I wasn’t really happy.”"
archive := "“I 😀 😬 😁 I 😂 🤣 😃 😄 😬 😅 😁 I 😆 I 😇 😉 😬, 😊 I 🙂’t 🙃 😬.”"
dictionary := '{"and":"😅","because":"😁","but":"😊","feel":"😉","felt":"😀","happy":"😬","knew":"😆","others":"😃","really":"🙃","saw":"😂","should":"😇","the":"🤣","wasn":"🙂","were":"😄"}'
```
See more [details here](../dictionary/README.md).

## Contributing
If you want us to add a new encoding algorithm, or you have found a bug or you just want to improve our project
please add an issue our github tracker.


## About
These project is built for fun and academic purposes in Go.

Copyright (c) 2017 [B.G.Adrian](https://coder.today) & @Davidescus
