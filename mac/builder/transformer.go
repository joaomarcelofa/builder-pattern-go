package builder

import (
	"fmt"
	"reflect"
)

func TransformMacInterface(macInterface basicInterface) (*MacBook, error) {
	macbook, ok := reflect.ValueOf(macInterface).Interface().(*MacBook)
	if !ok {
		return nil, fmt.Errorf("error at casting mac interface to macbook concrete struct")
	}
	return macbook, nil
}
