package main

import (
	_ "github.com/ludovicdeluna/go_tests/inspect"
	_ "github.com/ludovicdeluna/go_tests/samples/pointers"
	"github.com/ludovicdeluna/go_tests/ui.v2" // We can versionize package name with suffix.
)

func main() {
	ui.ClearScreen()
	samples := []string{"hello"}

	for stop := false; stop == false; {
		stop, _ = ui.ChooseSample(samples)
	}
}
