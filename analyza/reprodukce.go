package analyza

import (
	"math/rand"
	"strings"
)

const SLOVA = true
const ZNAKY = false

func (str *Text) Reproduct(what bool, combien, lenght uint64) string {
	dict := str.analyse(what, combien)
	return product(dict, what, combien, lenght)
}

func (str *Text) analyse(what bool, combien uint64) (dict map[string][]string) {
	dict = make(map[string][]string)
	if what == SLOVA {
		words := strings.Fields(string(*str))
		for i := combien; i < uint64(len(words)); i++ {
			var s string
			for j := uint64(1); j <= combien; j++ {
				s += words[i-j] + " "
			}
			dict[s] = append(dict[s], words[i])
		}
	} else {
		data := []rune(*str)
		for i := combien; i < uint64(len(data)); i++ {
			var c string
			for j := uint64(1); j <= combien; j++ {
				c += string(data[i-j])
			}
			dict[c] = append(dict[c], string(data[i]))
		}
	}
	return
}

func product(dict map[string][]string, what bool, combien uint64, lenght uint64) (out string) {
	var keys []string

	for k := range dict {
		keys = append(keys, k)
	}
	out = keys[rand.Intn(len(keys)-1)]
	if what == SLOVA {
		for i := combien; i < lenght; i++ {
			words := strings.Fields(out)
			var s string
			for j := combien; j > uint64(0); j-- {
				s += words[i-j]
			}
			if len(dict[s]) == 0 {
				out += keys[rand.Intn(len(keys)-1)]
			} else if len(dict[s]) == 1 {
				out += dict[s][0]
			} else {
				out += dict[s][rand.Intn(len(dict[s])-1)]
			}
		}
	} else {
		for i := combien; i < lenght; i++ {
			var c string
			for j := combien; j > uint64(0); j-- {
				c += string([]rune(out)[i-j])
			}
			if len(dict[c]) == 0 {
				out += keys[rand.Intn(len(keys)-1)]
			} else if len(dict[c]) == 1 {
				out += dict[c][0]
			} else {
				out += dict[c][rand.Intn(len(dict[c])-1)]
			}
		}
	}
	return
}
