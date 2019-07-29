package main

import "fmt"

func verbosef(format string, args ...interface{}) {
	if isVerbose {
		fmt.Printf(format+"\n", args...)
	}
}
