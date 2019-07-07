package analyza

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

type Text string

var samohlasky = []rune{'a', 'e', 'i', 'o', 'u'}
var souhlasky = []rune{'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'v', 'z'}
var slabtvrsouh = []rune{'r', 'l', 'm'} // slabikotvorné souhlásky
var konce = []rune{'.', '!', '?'}       // Konce vět
var whiteS = []rune{' ', '\n', '\t'}    // bílé znaky

func isIn(what rune, in []rune) bool {
	for _, s := range in {
		if what == s {
			return true
		}
	}
	return false
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func (in *Text) Stdr() (out Text) { // převede na málá písmena a pouze ASCII
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	tmp, _, _ := transform.String(t, string(*in))
	out = Text(tmp)
	out = Text(strings.ToLower(string(out)))
	return
}

func (str *Text) Frekvence() (fr map[rune]uint64) { // frekvence znaků
	fr = make(map[rune]uint64)
	for _, s := range *str {
		fr[s]++
	}
	return
}

func (str *Text) Words() (w uint64) { // Spočítá slova v textu. dle WhiteSpace
	var prevW bool = false
	for i, s := range *str {
		if isIn(s, whiteS) && prevW == true {
			w++
			prevW = false
		} else if !isIn(s, whiteS) {
			prevW = true
		}
		if prevW && i == len(*str)-1 {
			w++
		}
	}
	return
}

func (str Text) Slabiky() (sl uint64) {
	str = str.Stdr()
	return
}

func (str *Text) Sentences() (se uint64) { // spočítá věty v textu, končí .;?;! a následuje whiteSpace
	var prevE = false
	for i, s := range *str {
		if isIn(s, whiteS) && prevE {
			se++
		}
		prevE = false
		if isIn(s, konce) {
			prevE = true
		}
		if isIn(s, konce) && i == len(*str)-1 {
			se++
		}
	}
	return
}
