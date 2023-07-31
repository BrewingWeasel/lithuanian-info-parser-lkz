package parsing

import (
	"golang.org/x/exp/slices"
	"strings"
	"unicode"

	"github.com/charmbracelet/log"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var LITH_SPECIAL_CHARS = []rune{'ą', 'č', 'ę', 'ė', 'į', 'š', 'ų', 'ū', 'ž'}

func GetVerbInfo(grammaticalInfo string) [3]string {
	verbData := strings.Split(removeAccentuation(grammaticalInfo), ",")
	log.Debug(verbData)
	verbParts := [3]string{verbData[0], verbData[1], strings.SplitN(verbData[2], " ", 3)[1]}
	return createVerbVals(verbParts)

}

func getAccents(r rune) bool {
	if slices.Contains(LITH_SPECIAL_CHARS, r) {
		return false
	}
	return unicode.Is(unicode.Mn, r)
}

// Removes actual accentuation marks, not standard Lithuanian characters
// ie: pralaũkęs: pralaukęs
func removeAccentuation(word string) string {
	t := transform.Chain(runes.Remove(runes.Predicate(getAccents)), norm.NFC)
	result, _, _ := transform.String(t, word)
	return result
}

func createVerbVals(verbInfo [3]string) [3]string {
	log.Info(verbInfo)
	verbInfo[1] = strings.TrimSpace(verbInfo[1])
	verbInfo[2] = strings.TrimSpace(verbInfo[2])
	if strings.HasPrefix(verbInfo[1], "-") {
		log.Debug("Using prefixed")
		if len(verbInfo[1]) == 2 {
			removedVerbEnding := strings.TrimSuffix(verbInfo[0], "ti")
			if strings.HasSuffix(removedVerbEnding, "y") || strings.HasSuffix(removedVerbEnding, "ė") {
				removedVerbEnding = removedVerbEnding[:len(removedVerbEnding)-2]
			}
			verbInfo[1] = removedVerbEnding + verbInfo[1][1:]
			verbInfo[2] = removedVerbEnding + verbInfo[2][1:]
		}
		for i := len(verbInfo[0]) - 2; i > 0; i-- {
			if verbInfo[0][i:i+2] == verbInfo[1][1:3] {
				verbInfo[1] = verbInfo[0][:i] + verbInfo[1][1:]
				verbInfo[2] = verbInfo[0][:i] + verbInfo[2][1:]
				break
			}
		}
	}
	return verbInfo
}
