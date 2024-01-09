package main

import (
	"fmt"
	"github.com/skosovsky/go-head-first/ipkg"
)

func main() {
	value := ipkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
