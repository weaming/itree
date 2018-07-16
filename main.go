package main

import (
	"flag"

	ft "github.com/weaming/disk-analysis/filetree"
)

var (
	tree  = true
	level = 999
)

func init() {
	flag.BoolVar(&tree, "t", tree, "print in tree mode")
	flag.IntVar(&level, "L", level, "level to print in tree mode")
	flag.Parse()
}

func main() {
	path := "."
	root := ft.NewFileNode(path, path, nil)
	if tree {
		ft.PrintFileNodeTree(root, []string{}, 1, level)
	} else {
		ft.PrintFileNodeSimple(root)
	}
}
