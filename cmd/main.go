package main

import (
	"fmt"

	"github.com/azr/go-modules-example/subpkg/v2"
	gme "github.com/azr/go-modules-example/v2"
)

func main() {
	fmt.Println("root pkg version is ", gme.Version)
	fmt.Println("sub pkg version is ", subpkg.Version)
}
