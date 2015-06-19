package server

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
)

var (
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
)

func actionWrapper(action func(context *cli.Context) error) func(context *cli.Context) {
	return func(context *cli.Context) {
		if err := action(context); err != nil {
			logger.Fatalln(err)
		}
	}
}
