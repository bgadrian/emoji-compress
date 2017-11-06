package dictionary

import (
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bgadrian/emoji-compressor/emojis"

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
func TestAllForOne(t *testing.T) {
	source := ".! alfa /?'"
	emoji := ""
	var err error
	db := emojis.Database{}

	for err == nil {
		emoji, err = db.Fetch()
		if err != nil {
			if err.Error() != emojis.ErrFinished {
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
	table := []string{
		"",
		" ",
		"                              ",
		" `~!@#$%^&*()_+=-:\";'{}[]<>?/.,",
		// "😆~", //not supported yet
		"Alfa's arm is smaller then beta's,beta's arms is bigger!",
		"new .\n line %%",
		"One 25 year old twin stays on earth while the other, fresh out of astronaut school, sets off on a space voyage travelling at 90%% of the speed of light.\n After 10 years in space, with her mission accomplished, she turns round and heads back to earth. By the time she lands she knows from her on-board clock that 20 years have passed. She is now 45 years old. Fortunately, her study of relativity has prepared her for the shock when she sees her twin sister, who is now 71 years old.",
		//TODO add more crazy utf8 scenarios
	}

	for _, source := range table {
		comp, err := CompressString(source)

		if err != nil {
			t.Error(err)
			continue
		}

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
	t.Errorf("source malformed after decompress \nexp %s %s \n got %s", source, txtDiffs, decomp)

}
