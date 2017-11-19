package dictionary

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bgadrian/emoji-compress/emojis"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestCompressBasic(t *testing.T) {
	r, err := Compress([]byte("alfa beta teta alfa"))

	if err != nil {
		t.Error(err)
	}

	if len(r.Words) != 3 {
		t.Errorf("expected 3 words, found %v", len(r.Words))
	}

	c := utf8.RuneCountInString(r.Archive)

	if c != 7 {
		t.Errorf("archive result 7 chars, found %d : %s", c, r.Archive)
	}
}

func TestNotAllowed(t *testing.T) {
	_, err := CompressString("alfa 😅")
	if err == nil {
		t.Error("using an emoji shouldn't be allowed, yet")
	}
}

func TestDecompressBasic(t *testing.T) {
	w := map[string]string{
		"alfa": "🤣",
		"beta": "😇",
	}
	archive := "🤣 😇 🤣."
	source := "alfa beta alfa."

	result, err := DecompressString(w, archive)
	if err != nil {
		t.Error(err)
	}

	if result != source {
		t.Errorf("expected %v, got %v", source, result)
	}
}

//we test a decompress for each emoji in the DB
func TestDecompressAllEmojis(t *testing.T) {
	source := ".! alfa /?'"
	emoji := ""
	var err error
	db := emojis.Iterator{}

	for err == nil {
		emoji, err = db.NextSingleRune()
		if err != nil {
			if err != emojis.EOF {
				t.Error(err)
			}
			break
		}

		decomp, err := Decompress(&Result{
			Words: map[string]string{
				"alfa": emoji,
			},
			Archive: strings.Replace(source, "alfa", emoji, -1),
		})
		if err != nil {
			t.Error(err)
			continue
		}

		if strings.Compare(decomp, source) != 0 {
			t.Errorf("Decompress failed for %s %v", emoji, []rune(emoji))
		}
	}
}

func TestTable(t *testing.T) {
	words := []string{
		"",
		" `~!@#$%^&*",
		" `()_+=-:\";'",
		" `{}[]<>?/.,",
		// "😆~", //not supported yet
		" \n line %%",
		//from here https://golang.org/src/unicode/utf8/utf8_test.go
		"語þ日¥本¼語i日©",
		"日a本b語ç日ð本Ê",
		"日a本b語ç日ð本Ê語þ",
		"日¥本¼語i日©日a本b,",
		"語ç日ð本Ê語þ日¥本¼語",
		"i日©日a本b語ç日ð本Ê語þ日¥本¼語i日©",
	}

	phrases := []string{
		"%",
		"%, %.",
		"%, %? %...%! % .",
	}

	for _, word := range words {
		for _, phrase := range phrases {
			source := strings.Replace(phrase, "%", word, -1)
			comp, err := CompressString(source)

			if err != nil {
				t.Error(err)
				continue
			}
			// fmt.Println(comp.Archive)

			decomp, err := Decompress(comp)
			if err != nil {
				t.Error(err)
				continue
			}
			// fmt.Println(source, "=>", comp.Archive)

			if strings.Compare(decomp, source) != 0 {
				showDiff(source, decomp, t)
			}
		}
	}
}

func showDiff(source, decomp string, t *testing.T) {
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(source, decomp, false)
	txtDiffs := ""
	for _, diff := range diffs {
		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			txtDiffs += "\n++		" + diff.Text
		case diffmatchpatch.DiffDelete:
			txtDiffs += "\n--		" + diff.Text
		}
	}
	t.Errorf("source malformed after decompress \nexp %+q \n%+q \n got %+q", source, txtDiffs, decomp)

}

func Example() {
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
}

func ExampleDecompressString() {
	archive := "😀 😬 my 😁, my 😂, 🤣, 😃 😄 😬:" +
		"\n😅 😆 😇 😉 😊 🙂 😇 🙃 😋?" +
		"\nNo 😂, my 😂, 😌 😇 😍 😘 😂 😗—" +
		"\n😙 😚 😜 😝 😋 😇 🙃 😛 😊."

	dict := map[string]string{
		"All": "😙", "Take": "😀", "What": "😅", "all": "😬",
		"before": "😋", "call": "😗", "hadst": "🙃", "hast": "😆",
		"love": "😂", "loves": "😁", "mayst": "😍", "mine": "😚",
		"more": "😊", "take": "😃", "than": "🙂", "that": "😌",
		"them": "😄", "then": "😉", "thine": "😝", "this": "😛",
		"thou": "😇", "true": "😘", "was": "😜", "yea": "🤣",
	}

	original, err := DecompressString(dict, archive)

	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("Poetry: %s", original)
	// Output: Poetry: Take all my loves, my love, yea, take them all:
	// What hast thou then more than thou hadst before?
	// No love, my love, that thou mayst true love call—
	// All mine was thine before thou hadst this more.
}
