package main

import (
	"fmt"

	"github.com/alyumi/fury/test"
)

func main() {
	fmt.Println("Hello!")
	test.Test_perf(2, 5)
}
