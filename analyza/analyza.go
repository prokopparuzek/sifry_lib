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

func (str Text) Slabiky() (sl uint64) { // Spočítá slabiky
	str = str.Stdr()
	words := strings.Fields(string(str))
	println(words)
	for _, w := range words {
		for i, c := range w {
			switch {
			case isIn(c, samohlasky):
				if i < len(w)-1 {
					if isIn(rune(w[i+1]), samohlasky) && i != 2 && i != 3 {
					} else {
						sl++
					}
				} else {
					sl++
				}
			case isIn(c, slabtvrsouh):
				if i > 0 && i < len(w)-1 {
					if isIn(rune(w[i-1]), souhlasky) && isIn(rune(w[i+1]), souhlasky) {
						sl++
					}
				} else if i > 0 {
					if isIn(rune(w[i-1]), souhlasky) {
						sl++
					}
				} else if i < len(w)-1 {
					if isIn(rune(w[i+1]), souhlasky) {
						sl++
					}
				}
			}
		}
	}
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

func (str *Text) Flesh() float64 {
	var sl float64 = float64(str.Slabiky())
	var w float64 = float64(str.Words())
	var se float64 = float64(str.Sentences())
	return 206.835 - (1.015 * (w / se)) - 84.6*(sl/w)
}
