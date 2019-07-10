package change

import (
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"log"
	"regexp"
	"strings"
	"unicode"
)

func AlphaD(str *string) { // Odstraní vše co není číslo || písmeno(ASCII) || bílý znak
	reg, err := regexp.Compile("[^A-Za-z0-9 \t\n]+")
	if err != nil {
		log.Fatal("Nelze zkompilovat regex!")
	}
	*str = reg.ReplaceAllString(*str, "")
}

func RemoveWS(str *string) { // Odstraní bílé znaky
	*str = strings.Replace(*str, " ", "", -1)
	*str = strings.Replace(*str, "\t", "", -1)
	*str = strings.Replace(*str, "\n", "", -1)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

func Stdr(in *string) (out string) { // převede na malá písmena a pouze ASCII
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	tmp, _, _ := transform.String(t, *in)
	out = tmp
	out = strings.ToLower(out)
	return
}
