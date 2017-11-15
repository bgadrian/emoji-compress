/*Package dictionary is a small package that allows encoding (or compression) of strings  by replacing each unique word with an emoji.

Each compress generates a new dictionary and an encoded version of the text (archive),
based on the words found in the text.
If the original text doesn't have many repeating words, the "archive" will be longer
than the original string.

The dictionary should be sent to the user/client so he can decode the string.

Limitations:
* you cannot have emojis in the original text
* only works with a max of 1000 unique words
* (for now) compress generates a new dictionary for each text
* you have to use the same dictionary resulted from the Compress into the Decompress

The algorithm is very simple: tries to extract each word from the original text and replace it with an emoji.
A dictionary/map is generated along with the "Archive", to remember which word was replaced with each emoji.

The decompress process requires the "Archived" (encoded) version of the text, and the dictionary, in order to reverse the process.

TODO:
* the ability to use a custom dictionary when Compressing */
package dictionary

//Result contains the original phrase, dictionary and other data for Compress and Decompress
type Result struct {
	Words   map[string]string `json:"dict"`
	Source  string            `json:"source"`
	Archive string            `json:"archive"`
	Ratio   float32           `json:"ratio"`
}

//minimum amount of runes for a word to be compressed
const min int = 2
