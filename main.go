package main

import (
	"fmt"
	"github.com/utkarshg42/gostarter/stringutils"
)

func main() {
	a := "hi, gotest"
	fmt.Println(a, stringutils.Reverse(a))
}
