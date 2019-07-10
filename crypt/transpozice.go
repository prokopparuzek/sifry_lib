package crypt

import (
	"strings"
)

type Rectangle struct {
	Weight uint64
	Height uint64
}

func (r *Rectangle) CryptL(plain *string) (crypt string) {
	str := strings.Replace(*plain, "\n", "", -1)
	for i := uint64(0); i < r.Height; i++ {
		for j := uint64(0); j < r.Weight; j++ {
			if i*r.Weight+j < uint64(len(str)) {
				crypt += string([]rune(str)[i*r.Weight+j])
			} else {
				crypt += "X"
			}
		}
		crypt += "\n"
	}
	return
}

func (r *Rectangle) DecryptL(crypt *string) (plain string) {
	plain = strings.Replace(*crypt, "\n", "", -1)
	return
}

func (r *Rectangle) CryptR(plain *string) (crypt string) {
	str := []rune(*plain)
	for i := uint64(0); i < r.Height; i++ {
		for j := uint64(0); j < r.Weight; j++ {
			if j*r.Height+i < uint64(len(str)) {
				crypt += string(str[j*r.Height+i])
			} else {
				crypt += "X"
			}
		}
		crypt += "\n"
	}
	return
}

func (r *Rectangle) DecryptR(crypt *string) (plain string) {
	raw := strings.Replace(*crypt, "\n", "", -1)
	str := []rune(raw)
	for i := uint64(0); i < r.Height; i++ {
		for j := uint64(0); j < r.Weight; j++ {
			plain += string(str[j*r.Height+i])
		}
	}
	return
}
