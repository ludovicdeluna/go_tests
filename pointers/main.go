package main

import (
  "github.com/ludovicdeluna/go_tests/pointers/ui"
  _ "github.com/ludovicdeluna/go_tests/pointers/samples"
  _ "github.com/ludovicdeluna/go_tests/pointers/inspect"

)

func main () {
  ui.ClearScreen()
  samples := []string{"hello"}

  for stop := false ; stop == false ; {
    stop, _ = ui.ChooseSample(samples)
  }
}
