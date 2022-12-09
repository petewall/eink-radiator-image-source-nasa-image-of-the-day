// +build tools

package internal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	// These tools required by ginkgo
	_ "github.com/go-task/slim-sprig"
	_ "github.com/google/pprof/profile"
	_ "golang.org/x/tools/go/ast/inspector"

	// These tools required to run counterfeiter
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
)
