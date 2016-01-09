// This package will be ignored by "go install" or "go build" in any
// no Linux system. Thanks to "_linux.go" in the filename, automatically
// handled by the go build chain.
// See: https://golang.org/pkg/go/build/
package ui

import (
  "os"
  "os/exec"
)

func ClearScreen() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
