// This package will be ignored by "go install" or "go build" in any
// no Windows system. Thanks to "_windows.go" in the filename, automatically
// handled by the go build chain.
// See: https://golang.org/pkg/go/build/
package ui

import (
  "os"
  "os/exec"
)

func ClearScreen() {
	cmd := exec.Command("cls") //Windows example, not tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}
