package shell

import (
	"os"
	"os/exec"

	"github.com/tsaikd/KDGoLib/errutil"
)

// Option for Run
type Option struct {
	Args []string

	Append  bool
	Stdouts []string
	Stderrs []string
}

// Run shell command with Option
func Run(name string, option Option) (err error) {
	cmd := exec.Command(name, option.Args...)

	fileFlag := os.O_WRONLY
	if option.Append {
		fileFlag = fileFlag | os.O_APPEND
	} else {
		fileFlag = fileFlag | os.O_CREATE | os.O_TRUNC
	}
	filePerm := os.FileMode(0600)

	stdout := newBrocaster()
	stdout.AddOutput(os.Stdout)
	for _, name := range option.Stdouts {
		if err = stdout.AddFile(name, fileFlag, filePerm); err != nil {
			return
		}
	}
	cmd.Stdout = stdout
	defer func() {
		errutil.Trace(stdout.Close())
	}()

	stderr := newBrocaster()
	stderr.AddOutput(os.Stderr)
	for _, name := range option.Stderrs {
		if err = stderr.AddFile(name, fileFlag, filePerm); err != nil {
			return
		}
	}
	cmd.Stderr = stderr
	defer func() {
		errutil.Trace(stderr.Close())
	}()

	return cmd.Run()
}
