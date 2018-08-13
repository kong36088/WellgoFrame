package openapi

import (
	"github.com/kong36088/wellgo"
	"fmt"
	"github.com/kong36088/WellgoFrame/errors"
)

type Test struct {
	wellgo.Controller

	int   int     `param:"i",validate:"required"`
	float float64 `param:"float",validate:"required,gte=0"`
	str   string  `param:"str",validate:"required,"`
	arr   []interface{}
}

func (this *Test) Init(ctx *wellgo.WContext) {
	this.Controller.Init(ctx)

	fmt.Println("child init")
}

func (this *Test) Run() *wellgo.Result {
	return wellgo.NewResult(wellgo.GetErrorCode(errors.ErrTestErr), "", map[string]interface{}{"a": "b"})
}
