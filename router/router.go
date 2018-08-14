/**
  * @author wellsjiang
  * @date 2018/8/14
  */

package router

import (
	"github.com/kong36088/wellgo"
	"github.com/kong36088/WellgoFrame/controllers/openapi"
)

func InitRouter() {
	router := wellgo.GetRouterInstance()

	router.Register("app.api.test", &openapi.Test{})
	router.RegexpRegister("app\\.api\\..+", &openapi.Test{})
}
