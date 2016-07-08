package main

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"github.com/tsaikd/godson/server"
)

func main() {
	cmder.Main(
		*server.Module,
	)
}
