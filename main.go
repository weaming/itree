package main

import (
	"flag"

	ft "github.com/weaming/disk-analysis/filetree"
)

var (
	tree  = true
	human = true
	level = 1024
	path  = "."
)

func init() {
	flag.BoolVar(&tree, "t", tree, "tree mode")
	flag.BoolVar(&human, "h", human, "human readable size")
	flag.IntVar(&level, "L", level, "level in tree mode")
	flag.Parse()

	p := flag.Arg(0)
	if p != "" {
		path = p
	}
}

func main() {
	root := ft.NewFileNode(path, path, nil, true)
	if tree {
		ft.PrintFileNodeTree(root, []string{}, 1, level, human)
	} else {
		ft.PrintFileNodeSimple(root, human)
	}
}
