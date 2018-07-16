package main

import (
	"flag"

	ft "github.com/weaming/disk-analysis/filetree"
)

var (
	tree  = false
	human = false
	level = 999
)

func init() {
	flag.BoolVar(&tree, "t", tree, "print in tree mode")
	flag.BoolVar(&human, "human", human, "print in human readable size")
	flag.IntVar(&level, "L", level, "level to print in tree mode")
	flag.Parse()
}

func main() {
	path := "."
	root := ft.NewFileNode(path, path, nil)
	if tree {
		ft.PrintFileNodeTree(root, []string{}, 1, level, human)
	} else {
		ft.PrintFileNodeSimple(root, human)
	}
}
