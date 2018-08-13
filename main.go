package main

import (
	"github.com/kong36088/wellgo"
	"github.com/kong36088/WellgoFrame/controllers/openapi"
	"github.com/kong36088/WellgoFrame/errors"
	"os"
	"fmt"
)

func main() {

	workPath, _ := os.Getwd()
	fmt.Println(workPath)

	router := wellgo.GetRouterInstance()

	router.Register("app.openapi.test", &openapi.Test{})

	errors.InitErr()

	wellgo.Run()
}
