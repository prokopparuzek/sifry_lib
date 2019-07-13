package analyza

import (
	"prokop/sifry/change"
	"strings"
)

var samohlasky = []rune{'A', 'E', 'I', 'O', 'U'}
var souhlasky = []rune{'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'R', 'S', 'T', 'V', 'Z'}
var slabtvrsouh = []rune{'R', 'L'}   // slabikotvorné souhlásky
var konce = []rune{'.', '!', '?'}    // Konce vět
var whiteS = []rune{' ', '\n', '\t'} // bílé znaky

func isIn(what rune, in []rune) bool {
	for _, s := range in {
		if what == s {
			return true
		}
	}
	return false
}

func Frekvence(str *string) (fr map[string]uint64) { // frekvence znaků
	fr = make(map[string]uint64)
	for _, s := range *str {
		fr[string(s)]++
	}
	return
}

func FrekvenceSlov(str *string) (fr map[string]uint64) { // frekvence slov
	fr = make(map[string]uint64)
	w := strings.Fields(*str)
	for _, s := range w {
		fr[s]++
	}
	return
}

func Words(str *string) (w uint64) { // Spočítá slova v textu. dle WhiteSpace
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

func Slabiky(str string) (sl uint64) { // Spočítá slabiky
	str = change.Stdr(&str)
	words := strings.Fields(str)
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
			case c == 'm' && i == len(w)-1:
				sl++
			}
		}
	}
	return
}

func Sentences(str *string) (se uint64) { // spočítá věty v textu, končí .;?;! a následuje whiteSpace
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

func Flesh(str *string) float64 { // spočítá Fleshův index čitelnosti
	var sl float64 = float64(Slabiky(*str))
	var w float64 = float64(Words(str))
	var se float64 = float64(Sentences(str))
	return 206.835 - (1.015 * (w / se)) - 84.6*(sl/w)
}

func Lines(str *string) (ln uint64) { // Spočítá řádky.
	for _, c := range *str {
		if c == '\n' {
			ln++
		}
	}
	return
}

func Chars(str *string) uint64 { // Spočítá znaky.
	return uint64(len(*str))
}
