package emojis

import (
	"fmt"
	"testing"
)

func TestEmptyAndIsEmoji(t *testing.T) {
	db := Iterator{}
	emoji := ""
	var err error
	count := 0

	for err == nil {
		emoji, err = db.Next()
		if err != nil {
			if err != EOF {
				t.Error(err)
			}
			break
		}

		if len(emoji) == 0 {
			t.Errorf("found an empty emoji :( at %d runes %d", count, []rune(emoji))
		}

		if IsEmoji(emoji) == false {
			t.Errorf("isemoji didn't worked for %s", emoji)
		}

		count++
	}

	t.Logf("Found %d emojis", count)
}

func TestSingleRuneBasic(t *testing.T) {
	db := Iterator{}
	emoji := ""
	var err error
	count := 0

	for err == nil {
		emoji, err = db.NextSingleRune()
		if err != nil {
			if err != EOF {
				t.Error(err)
			}
			break
		}

		asRunes := []rune(emoji)
		if len(asRunes) != 1 {
			t.Errorf("singlerune() returned a multi rune %s %v", emoji, asRunes)
		}
		count++
	}
	t.Logf("Found %d single rune emojis", count)
}

func TestUnique(t *testing.T) {
	db := Iterator{}
	uniq := make(map[string]struct{})
	emoji := ""
	var err error
	count := 0

	for err == nil {
		emoji, err = db.Next()
		if err != nil {
			if err != EOF {
				t.Error(err)
			}
			break
		}

		count++
		_, ok := uniq[emoji]

		if ok {
			t.Errorf("emoji  %s is duplicate! at %d ", emoji, count)
		}

		uniq[emoji] = struct{}{}
	}
}

func Example() {
	db := &Iterator{}
	a, _ := db.Next()
	b, _ := db.Next()
	c, _ := db.Next()

	fmt.Printf("Emojis: %s %s %s", a, b, c)
	// Output: Emojis: 😀 😬 😁
}
