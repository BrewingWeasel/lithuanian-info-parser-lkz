package parsing

import "testing"

func TestRemoveAccentuation(t *testing.T) {
	result := removeAccentuation("supràsti")
	expected := "suprasti"

	if result != expected {
		t.Errorf("got %q, expected %q", result, expected)
	}
}
