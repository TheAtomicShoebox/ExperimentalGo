package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gitlab.com/schubachenterprises/experimentalgo/tree"
)

type Node = tree.Node

func main() {
	args := os.Args[1:]
	root := tree.CreateTree(0, convertArgs(args...))
	fmt.Printf("Root node: %v\n", &root)
}

func convertArgs(args ...string) map[int]int {
	mapArgs := make(map[int]int)
	for idx, arg := range args {
		if arg == "nil" {
			continue
		}

		val, err := strconv.Atoi(arg)

		if err != nil {
			log.Fatal(err)
		}

		mapArgs[idx] = val
	}
	return mapArgs
}
