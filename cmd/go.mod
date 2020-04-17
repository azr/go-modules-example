module github.com/azr/go-modules-example/cmd/v4

go 1.14

replace github.com/azr/go-modules-example/v4 => ../

replace github.com/azr/go-modules-example/subpkg/v4 => ../subpkg

require (
	github.com/azr/go-modules-example v0.0.0-20200319150029-d797a92ce60f
	github.com/azr/go-modules-example/subpkg v0.0.0-20200319150029-d797a92ce60f
	github.com/azr/go-modules-example/subpkg/v2 v2.0.0
	github.com/azr/go-modules-example/subpkg/v3 v3.0.0
	github.com/azr/go-modules-example/subpkg/v4 v4.0.0-00010101000000-000000000000
	github.com/azr/go-modules-example/v2 v2.0.0
	github.com/azr/go-modules-example/v3 v3.0.0
	github.com/azr/go-modules-example/v4 v4.0.0-00010101000000-000000000000
)
