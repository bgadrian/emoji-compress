package base64

import "testing"

func TestBasic(t *testing.T) {
	o, err := EncodeString("adasdbsad")

	if err != nil {
		t.Error(err)
	}

	output := string(o)
	// if output != "" {
	t.Errorf("exp %s got %s",
		"", output)
	// }
}
