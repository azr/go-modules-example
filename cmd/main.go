package main

import (
	"fmt"

	gmev1 "github.com/azr/go-modules-example"
	subpkgv1 "github.com/azr/go-modules-example/subpkg"

	subpkgv2 "github.com/azr/go-modules-example/subpkg/v2"
	gmev2 "github.com/azr/go-modules-example/v2"

	subpkgv3 "github.com/azr/go-modules-example/subpkg/v3"
	gmev3 "github.com/azr/go-modules-example/v3"
)

func main() {
	fmt.Println("gmev1 pkg version is ", gmev1.Version)
	fmt.Println("sub pkg version is ", subpkgv1.Version)

	fmt.Println("gmev2 pkg version is ", gmev2.Version)
	fmt.Println("subpkgv2 version is ", subpkgv2.Version)

	fmt.Println("gmev3 pkg version is ", gmev3.Version)
	fmt.Println("subpkgv3 version is ", subpkgv3.Version)
}
