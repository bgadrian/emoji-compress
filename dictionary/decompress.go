package dictionary

//DecompressString ...
func DecompressString(dict map[string]string, s string) (string, error) {
	return Decompress(&Result{
		Words:   dict,
		Archive: s,
	})
}

//Decompress Returns a string from the emoji encoded version.
func Decompress(r *Result) (string, error) {
	r.Source = ""
	reversed := make(map[string]string, len(r.Words))

	for s, r := range r.Words {
		reversed[r] = s
	}

	//go trough each character
	for _, nextRune := range r.Archive {
		word, ok := reversed[string(nextRune)]

		if ok {
			r.Source += word
			continue
		}

		r.Source += string(nextRune)
	}

	r.Ratio = float32(len(r.Archive)) / float32(len(r.Source))
	return r.Source, nil
}
