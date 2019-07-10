package transpozice

import (
	"strings"
)

type Rectangle struct {
	weight uint64
	height uint64
}

func (r *Rectangle) crypt(plain *string) (crypt string) {
	str := strings.Replace(*plain, "\n", "", -1)
	for i := uint64(0); i < r.height; i++ {
		for j := uint64(0); j < r.weight; j++ {
			if i*r.weight+j < uint64(len(str)) {
				crypt += string([]rune(str)[i*r.weight+j])
			} else {
				crypt += "x"
			}
		}
		crypt += "\n"
	}
	return
}

func (r *Rectangle) Decrypt(crypt *string) (plain string) {
	plain = strings.Replace(*crypt, "\n", "", -1)
	return
}
