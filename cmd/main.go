package main

import (
	"github.com/eatonphil/gosql"
	"github.com/eatonphil/gosql/LSM"
)

func main() {
	mb := LSM.NewMemoryBackend()

	gosql.RunRepl(mb)
}

//Everything in a folder should have the same package name except for main
