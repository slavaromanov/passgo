package passgo

import "os"

func main() {
	s := PassGen(os.Args[1:])
	print(compare(s, "\r\n\x00"))
}
