/**
 * @author wellsjiang
 */

package errors

import (
	"errors"
	"github.com/kong36088/wellgo"
)

var (
	ErrTestErr = errors.New("test error")

	ErrMap = map[error]int64{
		ErrTestErr: 10000,
	}
)

func InitErr() {
	wellgo.RegisterErrorMap(ErrMap)
}
