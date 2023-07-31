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
	"normalė́ti, -ė́ja, -ė́jo":   "normalėti, -ėja, -ėjo",
}

var getVerbFormsTests = map[[3]string][3]string{
	{"egzistuoti", "-uoja", "-avo"}: {"egzistuoti", "egzistuoja", "egzistavo"},
	{"sakyti", "sako", "sakė"}:      {"sakyti", "sako", "sakė"},
	{"barbenti", "-ena", "-eno"}:    {"barbenti", "barbena", "barbeno"},
	{"normalėti", "-ėja", "-ėjo"}:   {"normalėti", "normalėja", "normalėjo"},
	{"kaušti", "-ia", "-ė"}:         {"kaušti", "kaušia", "kaušė"},
	{"mušti", "-a", "-ė"}:           {"mušti", "muša", "mušė"},
}

func TestRemoveAccentuation(t *testing.T) {
	for testing, expected := range accentRemoval {
		result := removeAccentuation(testing)
		if result != expected {
			t.Errorf("got %q, expected %q", result, expected)
		}
	}
}

func TestGetVerbForms(t *testing.T) {
	for testing, expected := range getVerbFormsTests {
		result := createVerbVals(testing)
		if result != expected {
			t.Errorf("got %q, expected %q", result, expected)
		}
	}
}
