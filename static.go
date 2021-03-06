package passgo

var (
	l     int
	alpha = make(map[byte]bool, 128)
	// HelpMsg simple guide
	HelpMsg = "\nOptions:\n\t-d, --digits\tinclude digits\n\t-u, --upper \tinclude uppercase\n\t-l, --lower\tinclude lowercase\n\t-s, --spec\tinclude punctuation"
	getter  = map[byte][]byte{
		'u': []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		'l': []byte{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122},
		'd': []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57},
		's': []byte{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64, 91, 92, 93, 94, 95, 96, 123, 124, 125, 126},
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
