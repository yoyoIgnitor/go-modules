package main

import (
	"fmt"

	"github.com/yoyoIgnitor/go-modules/module2"
)

func main() {
	fmt.Println("Module 1 calling Module 222:")
	module2.DoSomething()
}
