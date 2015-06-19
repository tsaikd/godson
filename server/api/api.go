package api

import (
	"github.com/codegangsta/inject"
	"github.com/tsaikd/KDGoLib/errutil"
	"github.com/tsaikd/KDGoLib/injectutil"
)

var (
	apis []Handler
)

func Regist(handlers ...Handler) (err error) {
	if err = ValidateHandler(handlers...); err != nil {
		return
	}
	apis = append(apis, handlers...)
	return
}

// inject all registed API handler to inj
func Inject(inj inject.Injector) (err error) {
	for _, api := range apis {
		if _, err = injectutil.Invoke(inj, api); err != nil {
			return errutil.New("inject api handler failed", err)
		}
	}
	return
}
