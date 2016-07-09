package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/godson/server"
	"github.com/tsaikd/godson/shell"
)

func main() {
	cmder.Main(
		*server.Module,
		*shell.Module,
	)
}
