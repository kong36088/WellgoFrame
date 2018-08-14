package openapi

import (
	"github.com/kong36088/wellgo"
	"fmt"
	"github.com/kong36088/WellgoFrame/ferrors"
)

type Test struct {
	wellgo.Controller

	int   int     `param:"i",validate:"required"`
	float float64 `param:"float",validate:"required,gte=0"`
	str   string  `param:"str",validate:"required,"`
	arr   []interface{}
}

// 子构造函数，必须主动调用父构造函数
func (this *Test) Init(ctx *wellgo.WContext) {
	this.Controller.Init(ctx)

	fmt.Println("child init")
}

func (this *Test) Run() *wellgo.Result {
	this.Ctx.Logger.Debug(this.Args)
	this.Ctx.Logger.Debug(this)

	return wellgo.NewResult(wellgo.GetErrorCode(ferrors.ErrTestErr), "", map[string]interface{}{"a": "b"})
}
