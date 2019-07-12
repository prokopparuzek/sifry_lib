package crypt

import (
	"strings"
)

type Rectangle struct {
	Weight uint64
	Height uint64
}

type Teeth uint64
type Stairs uint64

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
	str := []rune(strings.Replace(*plain, "\n", "", -1))
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
	str := []rune(strings.Replace(*crypt, "\n", "", -1))
	for i := uint64(0); i < r.Height; i++ {
		for j := uint64(0); j < r.Weight; j++ {
			plain += string(str[j*r.Height+i])
		}
	}
	return
}

func Reverse(in *string) (out string) {
	str := []rune(*in)
	for i := len(str) - 1; i >= 0; i-- {
		out += string(str[i])
	}
	return
}

func Triangle90(in *string) (out string) {
	str := []rune(*in)
	index := 0
	for i := 1; ; i++ {
		for j := 0; j < i; j++ {
			if index >= len(str) {
				index++
				out += "X"
			} else {
				out += string(str[index])
				index++
			}
		}
		out += "\n"
		if index >= len(str) {
			break
		}
	}
	return
}

/* TODO
func Triangle(in *string) (out string) {

	return
}*/

func TriangleD(in *string) (out string) {
	out = strings.Replace(*in, "\n", "", -1)
	out = strings.Replace(*in, " ", "", -1)
	return
}

func (weight Stairs) Crypt(in *string) (out string) {
	index := 0
	str := []rune(*in)
	for i := 0; ; i++ {
		for j := 0; j < i*int(weight); j++ {
			out += " "
		}
		for j := 0; j < int(weight); j++ {
			if index == len(str) {
				out += "X"
			} else {
				out += string(str[index])
				index++
			}
		}
		out += "\n"
	}
	return
}

func (height Teeth) Crypt(in *string) (out string) {
	*in = strings.Replace(*in, "\n", "", -1)
	crypt := make([][]rune, height)
	index := 0
	dir := true
	for i := 0; i < int(height); i++ {
		crypt[i] = make([]rune, len(*in))
	}
	for i, c := range *in {
		if dir {
			crypt[index][i] = c
			index++
			if index == int(height) {
				dir = false
				index -= 2
			}
		} else {
			crypt[index][i] = c
			index--
			if index == -1 {
				dir = true
				index = 1
			}
		}
	}
	for _, ar := range crypt {
		for _, c := range ar {
			if c == 0 {
				out += " "
			} else {
				out += string(c)
			}
		}
		out += "\n"
	}
	return
}

func (height Teeth) Decrypt(in *string) (out string) {
	index := 0
	dir := true
	crypt := make([][]rune, height)
	tmp := strings.Split(*in, "\n")
	for i, s := range tmp {
		crypt[i] = []rune(s)
	}
	for i := 0; i < len(tmp[0]); i++ {
		if dir {
			out += string(crypt[index][i])
			index++
			if index == int(height) {
				dir = false
				index -= 2
			}
		} else {
			out += string(crypt[index][i])
			index--
			if index == -1 {
				dir = true
				index = 1
			}
		}
	}
	return
}
