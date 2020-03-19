package main

import (
	"fmt"

	gme "github.com/azr/go-modules-example"
	"github.com/azr/go-modules-example/subpkg"
)

func main() {
	fmt.Println("root pkg version is ", gme.Version)
	fmt.Println("sub pkg version is ", subpkg.Version)
}
