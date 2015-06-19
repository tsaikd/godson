package api

import (
	"errors"
	"reflect"
)

type Handler interface{}

func ValidateHandler(handlers ...Handler) (err error) {
	for _, handler := range handlers {
		if handler == nil || reflect.TypeOf(handler).Kind() != reflect.Func {
			err = errors.New("handler must be a callable func")
			return
		}
	}
	return
}
