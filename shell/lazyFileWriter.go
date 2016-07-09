package shell

import "os"

func newLazyFileWriter(name string, flag int, perm os.FileMode) *lazyFileWriter {
	if flag&os.O_CREATE != 0 {
		if err := os.Remove(name); err != nil {
			if !os.IsNotExist(err) {
				panic(err)
			}
		}
	}
	return &lazyFileWriter{
		name: name,
		flag: flag,
		perm: perm,
	}
}

type lazyFileWriter struct {
	name string
	flag int
	perm os.FileMode

	file *os.File
}

func (t *lazyFileWriter) Write(p []byte) (n int, err error) {
	if t.file == nil {
		if t.file, err = os.OpenFile(t.name, t.flag, t.perm); err != nil {
			return
		}
	}
	return t.file.Write(p)
}

func (t lazyFileWriter) Close() error {
	if t.file == nil {
		return nil
	}
	return t.file.Close()
}
