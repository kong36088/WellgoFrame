package main

import (
	"github.com/kong36088/wellgo"
	"github.com/kong36088/WellgoFrame/ferrors"
	"github.com/kong36088/WellgoFrame/router"
)

func main() {
	ferrors.InitErr()

	router.InitRouter()

	wellgo.Run()
}
