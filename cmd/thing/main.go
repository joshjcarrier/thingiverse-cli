package main

import (
	"github.com/joshjcarrier/thingiverse-cli/pkg/cmd"
)

var (
	version string // set by linker -X
)

func main() {
	cmd.Execute(version)
}
