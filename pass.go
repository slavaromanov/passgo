package passgo

import "os"

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

func lcg() func() uint32 {
	r := uint32(os.Getppid() + os.Getpid())
	return func() uint32 {
		r = (1103515245*r + 12345) % uint32((1<<31)-1)
		return r
	}
}

func pow10(pow int) int {
	res, val := 1, 10
	for pow > 0 {
		if pow&1 != 0 {
			res *= val
		}
		val *= val
		pow >>= 1 // equal `pow /= 2`
	}
	return res
}

func atoi(digit string) (r int) {
	p := pow10(len(digit) - 1)
	for _, c := range []byte(digit) {
		r += atoiMap[c] * p
		p /= 10
	}
	return
}

func randint(foo func() uint32, max int) func() int {
	return func() int {
		return int(foo()) % max
	}
}

func switchFilter(c byte) bool {
	b, ok := getter[c]
	if ok {
		for _, q := range b {
			alpha[q] = true
		}
	}
	return ok
}

func serializer(obj map[byte]bool) []byte {
	var b []byte
	for c := range obj {
		b = append(b, c)
	}
	return b
}

func serializeMap() map[bool][]byte {
	return map[bool][]byte{
		false: append(getter['l'], append(getter['u'], getter['d']...)...),
		true:  serializer(alpha),
	}
}

// GetAlphabet generate alphabeth by flags
func GetAlphabet(flags []string) []byte {
	for _, a := range flags {
		charr := []byte(a)
		if charr[0] == '-' {
			if all(string(charr[0:2]), []byte("-")) {
				switchFilter(charr[2])
			} else {
				for _, c := range charr[1:] {
					switchFilter(c)
				}
			}
		} else if all(a, getter['d']) {
			l = atoi(a)
		}
	}
	m := serializeMap()
	res := m[len(alpha) != 0]
	return res
}

// PassGen ... READ FLAGS AND GENERATING PASS
func PassGen(abc []byte) (pass []byte) {
	rnd := randint(lcg(), len(abc))
	for ; l > 0; l-- {
		pass = append(pass, abc[rnd()])
	}
	return pass
}

// NewPass generate password
func NewPass(fl []string) string {
	return string(PassGen(GetAlphabet(fl)))
}
