package main

import (
	"os"
)

var (
	err    error
	l      int
	alpha  []byte
	seedL  = 2
	e      = false
	filter = make([]bool, 4)
	m      = map[byte]int{'u': 0, 'l': 1, 'd': 2, 's': 3}
	help   = "\n\nOptions:\n\t-d, --digits\tinclude digits\n\t-u, --upper \tinclude uppercase\n\t-l, --lower\tinclude lowercase\n\t-s, --spec\tinclude punctuation"
	getter = map[int][]byte{
		0: []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		1: []byte{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122},
		2: []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57},
		3: []byte{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64, 91, 92, 93, 94, 95, 96, 123, 124, 125, 126},
	}
	atoiMap = map[byte]int{
		0x30: 0, 0x31: 1,
		0x32: 2, 0x33: 3,
		0x34: 4, 0x35: 5,
		0x36: 6, 0x37: 7,
		0x38: 8, 0x39: 9,
		0x61: 10, 0x62: 11,
		0x63: 12, 0x64: 13,
		0x65: 14, 0x66: 15,
	}
)

func any(char byte, set []byte) bool {
	for _, c := range set {
		if char == c {
			return true
		}
	}
	return false
}

func all(word string, set []byte) bool {
	for _, char := range []byte(word) {
		if !any(char, set) {
			return false
		}
	}
	return true
}

func lcg(a, c, m, seed uint32) func() uint32 {
	r := seed
	return func() uint32 {
		r = (a*r + c) % m
		return r
	}
}

func pow10(b int) int {
	p, a := 1, 10
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func atoi(digit string) int {
	p := pow10(len(digit) - 1)
	r := 0
	for _, c := range []byte(digit) {
		r += atoiMap[c] * p
		p /= 10
	}
	return r
}

func genSeed() uint32 {
	res := 0
	if fd, err := os.Open("/proc/sys/kernel/random/uuid"); err != nil {
		panic(err)
	} else {
		buff := make([]byte, seedL)
		if n, err := fd.Read(buff); err != nil {
			panic(err)
		} else {
			if n != seedL {
				println("READ FROM FILE", n)
			}
			for i, c := range buff {
				res += atoiMap[c] * (1 << uint(4*(seedL-1-i)))
			}
		}
	}
	return uint32(res)
}

func randint(foo func() uint32, max int) func() int {
	return func() int {
		return int(foo()) % max
	}
}

func main() {
	out := "Usage: \n\t" + os.Args[0] + " -dlus 32\n\ts:bXxv`kOZesS)Ll" + help
	if len(os.Args) > 1 {
		for _, a := range os.Args[1:] {
			if all(a, getter[2]) {
				l = atoi(a)
			} else {
				charr := []byte(a)
				if charr[0] != '-' {
					l, out = 0, ("Unknown argument " + a + help)
					break
				}
				if charr[1] == '-' {
					if i, ok := m[charr[2]]; ok {
						filter[i] = true
					} else {
						l, out = 0, ("Unknown argument " + a + help)
						break
					}
				} else {
					for _, c := range charr[1:] {
						if _, ok := m[c]; ok {
							filter[m[c]] = true
						} else {
							l, out = 0, ("Unknown option -" + string(c) + " in " + a + help)
							break
						}
					}
				}
			}
		}
		if l > 0 {
			out = ""
			for i, b := range filter {
				if b {
					alpha = append(alpha, getter[i]...)
				}
			}
			if len(alpha) == 0 {
				alpha = append(alpha, append(getter[0], append(getter[1], getter[2]...)...)...)
			}
			rnd := randint(lcg(1103515245, 12345, uint32((1<<31)-1), genSeed()), len(alpha))
			for ; l > 0; l-- {
				out += string(alpha[rnd()])
			}
		}
	}
	println(out + "\r\n\x00")
}
