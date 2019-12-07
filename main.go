package main

import (
	"os"

	"github.com/zabio3/spokesman/cmd"
)

func main() {
	cli := &cmd.CLI{ErrStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
