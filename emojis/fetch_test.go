package emojis

import (
	"testing"
)

func TestBasic(t *testing.T) {
	db := Database{}
	for i := 0; i < 13; i++ {
		emoji, err := db.Fetch()

		if err != nil {
			t.Error(err)
		}

		if len(emoji) == 0 || emoji == "" {
			t.Error("expected an emoji got empty string")
		}

		t.Log("got:"+emoji, []rune(emoji))
	}
}

func TestEmptyAndIsEmoji(t *testing.T) {
	db := Database{}
	emoji := ""
	var err error
	count := 0

	for err == nil {
		emoji, err = db.Fetch()
		if err != nil {
			if err.Error() != ErrFinished {
				t.Error(err)
			}
			break
		}

		if len(emoji) == 0 {
			t.Errorf("found an empty emoji :( at %d runes %d", count, []rune(emoji))
		}

		asRunes := []rune(emoji)
		if IsEmoji(asRunes[0]) == false {
			t.Errorf("illegal detection didn't worked for %s %v", emoji, asRunes)
		}
		count++
	}
}

func TestUnique(t *testing.T) {
	db := Database{}
	uniq := make(map[string]struct{}, 1000)
	emoji := ""
	var err error
	count := 0

	for err == nil {
		emoji, err = db.Fetch()
		if err != nil {
			if err.Error() != ErrFinished {
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
