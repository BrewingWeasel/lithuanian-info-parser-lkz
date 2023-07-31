package parsing

import "testing"

var accentRemoval = map[string]string{
	"supràsti":                  "suprasti",
	"suprañta":                  "supranta",
	"suprañtančiai":             "suprantančiai",
	"išpràsti":                  "išprasti",
	"svei̇̃kas":                  "sveikas",
	"labas":                      "labas",
	"egzistúoti, -úoja, -ãvo": "egzistuoti, -uoja, -avo",
}

func TestRemoveAccentuation(t *testing.T) {
	for testing, expected := range accentRemoval {
		result := removeAccentuation(testing)
		if result != expected {
			t.Errorf("got %q, expected %q", result, expected)
		}
	}
}
