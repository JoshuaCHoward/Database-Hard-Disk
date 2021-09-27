package main

import (
	gosql "home"
	"home/LSM"
)

func main() {
	mb := LSM.NewMemoryBackend()

	gosql.RunRepl(mb)
}

//Everything in a folder should have the same package name except for main
