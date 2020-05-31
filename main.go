package main

import (
	"fmt"
	"github.com/utkarshg42/gostarter/stringutils"
)

func main() {
	a := "hi, gotest"
	fmt.Println(a, stringutils.Reverse(a))
	fmt.Println(a, stringutils.Encode(a, 14), stringutils.Decode(stringutils.Encode(a, 14), 14))
}
