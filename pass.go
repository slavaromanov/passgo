package main

import (
	"math/rand"
	"time"

	"github.com/mkideal/cli"
	"github.com/atotto/clipboard"
)

type CharRange struct {
	min, max int
}



var uppers = CharRange{65,90}
var lowers = CharRange{97,122}
var digits = CharRange{49,56}
var special = []CharRange{
	CharRange{32,47},
	CharRange{58,64},
	CharRange{91,96},
	CharRange{123,126},
}

func makeRange(cr CharRange) []int {
	a := make([]int, cr.max-cr.min+1)
	for i := range a {
		a[i] = cr.min + i
	}
	return a
}


func makeSequence(cra []CharRange) []int {
	var res []int
	for _, cr := range cra {
		a := makeRange(cr)
		res = append(res, a[:]...)
	}
	return res
}

func rs(chars []int, l int, f bool) string {
	if !f {
		rand.Seed(time.Now().UTC().UnixNano())
	}
	result := ""
	for i := 0; i < l; i++ {
		if f {
			rand.Seed(time.Now().UTC().UnixNano())
		}
		index := rand.Intn(len(chars))
		result += string(chars[index])
	}
	return result
}


type argT struct {
    cli.Helper
    L	 int  `cli:"*l,length" usage:"define string length"`
    I    bool `cli:"i,ignore-case" usage:"boolean type ignore case register (lower only)"`
    S    bool `cli:"s,specials" usage:"boolean type include special characters"`
    C	 bool `cli:"c,clip" usage:"copy output to clipboards"`
    F    bool `cli:"!f,force" usage:"add more colision to GPRD"`
    Q	 bool `cli:"q,quiet" usage:"not output volume to stdout"`
}


func main() {
	cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		chars := []CharRange{lowers, digits}
		if !argv.I {
			chars = append(chars, uppers)
		}
		if argv.S == true {
			for _, x := range special {
				chars = append(chars, x)
			}
		}
		res := rs(makeSequence(chars), argv.L, argv.F)
		if argv.C == true {
			clipboard.WriteAll(res)
		}
		print(res, "\n")
		return nil
	})
}
