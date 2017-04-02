package main

import (
	"os"

	"github.com/slavaromanov/passgo"
)

func main() {
	flags := os.Args
	if len(flags) < 2 {
		println(passgo.HelpMsg)
		os.Exit(1)
	}
	println(passgo.NewPass(flags[1:]))
}
