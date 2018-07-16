package main

import (
	"flag"
	"os"

	ft "github.com/weaming/disk-analysis/filetree"
)

var (
	tree  = false
	human = false
	level = 1024
	path  = "."
)

func init() {
	flag.BoolVar(&tree, "t", tree, "tree mode")
	flag.BoolVar(&human, "human", human, "human readable size")
	flag.IntVar(&level, "L", level, "level in tree mode")
	flag.Parse()

	if flag.NArg() == 0 {
		println("missing positional argument PATH")
		os.Exit(1)
	}
	path = flag.Arg(0)
}

func main() {
	root := ft.NewFileNode(path, path, nil)
	if tree {
		ft.PrintFileNodeTree(root, []string{}, 1, level, human)
	} else {
		ft.PrintFileNodeSimple(root, human)
	}
}
