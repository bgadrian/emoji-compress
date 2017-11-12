package lz78

import "testing"
import "strings"

import "github.com/sergi/go-diff/diffmatchpatch"
import "fmt"
import "io/ioutil"

type casesInOut []inOut

type inOut struct {
	in  string
	out string
}

func TestDecompresInOut(t *testing.T) {

	c := casesInOut{
		{
			in:  "|0000a||0000b||0001b||0000c||0002a|",
			out: "ababcba",
		},
		{
			in:  "|0000a||0000b||0001b||0003c|",
			out: "abababc",
		},
		{
			in:  "|0000|||0000b||0001b||0003c|",
			out: "|b|b|bc",
		},
	}

	for _, e := range c {
		r, err := Decompress(e.in)
		if err != nil {
			t.Error(err)
		}

		if r != e.out {
			t.Errorf("Expected %s, got %s", e.out, r)
		}
	}
}

func TestFullTableASCII(t *testing.T) {
	table := []string{
		"",
		" ",
		"123456789012",
		"~!@~!@~!@~!@~!@~!@",
		"Broasca are sau nu are mere?",
		//utf8 diacritics
		"A fost odată ca-n poveşti,\nA fost ca niciodată.\nDin rude mari împărăteşti,\nO prea frumoasă fată.",
		"[原文]篭毛與 美篭母乳 布久思毛與 美夫君志持 此岳尓 菜採須兒 家吉閑名 告<紗>根 虚見津 山跡乃國者 押奈戸手 吾許曽居 師<吉>名倍手 吾己曽座 我<許>背齒 告目 ...",
		"Text^=+Text^=+Text^=+Tex^=+Tex^=+Tex^^=+Tx^^=+x^^=+x^^=+x^^=++x^^=+++x^^=+++x^~,!?:;'\"'`.%&*)([]{}|",
		" يرتبط القادة باستخدام. بالرغم واشتدّت باستخدام تعد من. كل الأرض لليابان ارتكبها لها, إذ بين ووصف تكتيكاً الإحتفاظ, من أضف ",
		"φαβελλασ πετεντιθμ vελ νε, ατ νισλ σονετ οπορτερε εθμ",
		//TODO add more here
	}

	for _, s := range table {
		err := compareInOut(s)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFullDynamicASCII(t *testing.T) {
	from := 0
	to := 128
	c := []string{
		"a", "aa", "a?", "aa?",
		"aba", "abba", "bbaa",
		//TODO add more here
	}

Exit:
	for _, chars := range c {
		for i := from; i < to; i++ {
			source := strings.Repeat(chars, i)
			err := compareInOut(source)

			if err != nil {
				t.Error(err)
				break Exit
			}
		}
	}
}

func TestHeavy(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Error("Recovered in TestHeavy", r)
		}
		return
	}()
	s, e := ioutil.ReadFile("big-poetry.txt")
	if e != nil {
		t.Error("Cant read file")
	}

	err := compareInOut(string(s))
	if err != nil {
		t.Error(err)
	}
}

func compareInOut(s string) error {
	c, err := CompressString(s)
	if err != nil {
		return err
	}

	d, err := Decompress(c)
	if err != nil {
		return err
	}

	if strings.Compare(s, d) != 0 {
		//because the texts are "aaabbababaaaba" it's very hard to see
		//the diff between source and output, so we're using a diff colored result
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(s, d, false)
		return fmt.Errorf("source malformed after decompress \nexp: %s \ndiff: %s \ngot: %s",
			s, dmp.DiffPrettyText(diffs), d)
	}

	return nil
}
