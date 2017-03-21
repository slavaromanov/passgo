package main

import (
  "os"
  "time"
  "regexp"
  "strconv"
  "math/rand"
)

var (
  err error
  length int
  alpha []byte
  filter = make([]bool, 4)
  m = map[byte]int{'u': 0, 'l': 1, 'd': 2, 's': 3}
  dre = regexp.MustCompile("^[0-9]*$")
  sre = regexp.MustCompile("^-([ulds]{1,4})$")
  lre = regexp.MustCompile("^--(upper|lower|spec|digits)$")
  help = "\n\nOptions:\n\t-d, --digits\tinclude digits\n\t-u, --upper \tinclude uppercase\n\t-l, --lower\tinclude lowercase\n\t-s, --spec\tinclude punctuation"
  getter = map[int][]byte{
    0: []byte{65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85,86, 87, 88, 89, 90},
    1: []byte{97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122},
    2: []byte{48, 49, 50, 51, 52, 53, 54, 55, 56, 57},
    3: []byte{33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 58, 59, 60, 61, 62, 63, 64, 91, 92,93, 94, 95, 96, 123, 124, 125, 126},
  }
)

func main() {
  out := "Usage: \n\t" + os.Args[0] + " -dlus 32\n\ts:bXxv`kOZesS)Ll" + help
  if len(os.Args) > 1 {
    rand.Seed(time.Now().UTC().UnixNano())
    for _, a := range os.Args[1:] {
      if sre.MatchString(a) {
        for _, c := range []byte(sre.FindStringSubmatch(a)[1]) {
          filter[m[c]] = true
        }
        } else if dre.MatchString(a) {
          length, err = strconv.Atoi(a)
          if err != nil {panic(err)}
        } else if lre.MatchString(a) {
            filter[m[[]byte(lre.FindStringSubmatch(a)[1])[0]]] = true
        } else {
          length, out = 0, ("Unknown argument " + a + help)
          break
        }
    }
    if length > 0 {
      out = ""
      for i, b := range filter {if b {alpha = append(alpha, getter[i]...)}}
      if len(alpha) == 0 {alpha = append(alpha, append(getter[0], append(getter[1], getter[2]...)...)...)}
      for i:=0; i<length; i++ {out += string(alpha[rand.Intn(len(alpha))])}
    }
  }
  println(out)
}
