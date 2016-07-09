package shell

import (
	"github.com/tsaikd/KDGoLib/cliutil/cmder"
	"gopkg.in/urfave/cli.v2"
)

// Module info
var Module = cmder.NewModule("shell").
	SetUsage("run shell script and log stdout/stderr without tee").
	AddFlag(
		&cli.BoolFlag{
			Name:        "a",
			Aliases:     []string{"append"},
			Usage:       "Append the output to the files rather than overwriting them",
			Destination: &flagAppend,
		},
		&cli.StringSliceFlag{
			Name:  "stdout",
			Usage: "Log stdout stream to file, could use multiple times",
			Value: &flagStdouts,
		},
		&cli.StringSliceFlag{
			Name:  "stderr",
			Usage: "Log stderr stream to file, could use multiple times",
			Value: &flagStderrs,
		},
	).
	SetAction(action)

var flagAppend bool
var flagStdouts = cli.StringSlice{}
var flagStderrs = cli.StringSlice{}

func action(c *cli.Context) (err error) {
	args := c.Args()
	return Run(args.First(), Option{
		Args:    args.Tail(),
		Append:  flagAppend,
		Stdouts: flagStdouts.Value(),
		Stderrs: flagStderrs.Value(),
	})
}
