package transpozice

import (
	"strings"
)

type Text string

func (plain *Text) RectangleC(w, h uint64) (crypt Text) {
	for i := uint64(0); i < h; i++ {
		for j := uint64(0); j < w; j++ {
			if i*w+j < uint64(len(*plain)) {
				crypt += Text([]rune(*plain)[i*w+j])
			} else {
				crypt += "x"
			}
		}
		crypt += "\n"
	}
	return
}

func (crypt *Text) RectangleD() (plain Text) {
	plain = Text(strings.Replace(string(*crypt), "\n", "", -1))
	return
}
