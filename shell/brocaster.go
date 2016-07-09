package shell

import (
	"io"
	"os"

	"github.com/tsaikd/KDGoLib/errutil"
)

// errors
var (
	ErrorNoOutputInBrocaster  = errutil.NewFactory("no output in brocaster")
	ErrorOutputCountMismatch2 = errutil.NewFactory("output count expect %d but got %d")
)

func newBrocaster() *brocaster {
	return &brocaster{
		outputs:    []io.Writer{},
		closeFuncs: []closeFunc{},
	}
}

type brocaster struct {
	outputs    []io.Writer
	closeFuncs []closeFunc
}

type closeFunc func() (err error)

var _ io.WriteCloser = brocaster{}

func (t brocaster) Write(p []byte) (n int, err error) {
	var errs = []error{}
	var ns = []int{}

	if len(t.outputs) < 1 {
		return 0, ErrorNoOutputInBrocaster.New(nil)
	}

	for _, output := range t.outputs {
		bn, berr := output.Write(p)
		if berr != nil {
			errs = append(errs, berr)
		} else {
			ns = append(ns, bn)
		}
	}

	// check output result
	if len(ns) > 0 {
		n = ns[0]
		for _, bn := range ns[1:] {
			if n != bn {
				return 0, ErrorOutputCountMismatch2.New(nil, n, bn)
			}
		}
	}

	if len(errs) > 0 {
		return n, errutil.NewErrors(errs...)
	}

	return n, nil
}

func (t brocaster) Close() error {
	var errs = []error{}

	for _, closer := range t.closeFuncs {
		if err := closer(); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errutil.NewErrors(errs...)
	}

	return nil
}

func (t *brocaster) AddOutput(output io.Writer) {
	t.outputs = append(t.outputs, output)
}

func (t *brocaster) AddCloseFunc(closer closeFunc) {
	t.closeFuncs = append(t.closeFuncs, closer)
}

func (t *brocaster) AddFile(name string, flag int, perm os.FileMode) (err error) {
	cache := brocasterFileCache[name]
	if cache != nil {
		t.AddOutput(cache)
		return nil
	}

	file := newLazyFileWriter(name, flag, perm)
	t.AddOutput(file)
	t.AddCloseFunc(file.Close)
	brocasterFileCache[name] = file
	return nil
}

var brocasterFileCache = map[string]*lazyFileWriter{}
