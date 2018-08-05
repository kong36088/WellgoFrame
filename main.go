package main

import (
	"fmt"
	"reflect"
)

type A struct{

}

type B struct{

}

func anoy(r interface{}){
	fmt.Print(reflect.TypeOf(r))
}

func main(){
	b := &B{}
	anoy(b)
}
