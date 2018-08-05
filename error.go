/**
 * @author wellsjiang
 */

package main

import (
	"errors"
	"github.com/kong36088/wellgo"
)

var (
	ErrTestErr = errors.New("test error")

	ErrMap = map[error]int{
		ErrTestErr: 10000,
	}
)

func InitErr() {
	wellgo.RegisterErrorMap(ErrMap)
}
