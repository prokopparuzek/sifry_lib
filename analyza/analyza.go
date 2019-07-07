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
var slabtvrsouh = []rune{'r', 'l', 'm'}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func (in Text) Stdr() (out Text) {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	tmp, _, _ := transform.String(t, string(in))
	out = Text(tmp)
	out = Text(strings.ToLower(string(out)))
	return
}

func (str Text) Frekvence() (fr map[rune]uint64) { // frekvence znaků
	fr = make(map[rune]uint64)
	for _, s := range str {
		fr[s]++
	}
	return
}

func (str Text) Words() (w uint64) { // Spočítá slova v textu.
	var prevW bool = false
	for _, s := range str {
		if (s == ' ' || s == '\t' || s == '\n') && prevW == true {
			prevW = false
		} else if s != ' ' && s != '\t' && s != '\n' && prevW == false {
			w++
			prevW = true
		}
	}
	return
}

func (str Text) Slabiky() (sl uint64) {
	str = str.Stdr()
	return
}
