# go-modules-example
trying things around go modules multimodules, multisubmodules and tagging

https://github.com/azr/go-modules-example/blob/v3/cmd/main.go ( in the `v3` branch ) uses modules of the same packages under 3 different versions.

here are some take aways :

* to tag a subpkg one just has to git tag `subpkg/v3.0.0` for example allowing to selectively tag subpackages like here: https://github.com/azr/go-modules-example/tree/subpkg/v3.0.0

* use a `v${MAJOR_VERSION}` branch to make updates easier, my go (`1.14`) could only find the `v3.0.0`  version - without hinting nor using the commit sha - only after I put the tag in the `v3` branch ( first it was in the `v2`  one )

* there always needs to be a `go.mod` file in the root git folder 
