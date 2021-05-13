package main

import (
	"flag"
	"fmt"

	ft "github.com/weaming/itree/filetree"
)

var (
	tree  = true
	human = true
	md5   = false
	imo   = false
	level = 1024
	paths = []string{}
)

func init() {
	flag.BoolVar(&tree, "t", tree, "tree mode")
	flag.BoolVar(&human, "h", human, "human readable size")
	flag.BoolVar(&md5, "md5", md5, "with MD5")
	flag.BoolVar(&imo, "imo", imo, "with ImoHash https://github.com/kalafut/imohash")
	flag.IntVar(&level, "L", level, "level in tree mode")
	flag.Parse()

	for i := 0; i < 1000; i++ {
		p := flag.Arg(i)
		if p != "" {
			paths = append(paths, p)
		} else {
			if i == 0 {
				paths = append(paths, ".")
			} else {
				break
			}
		}
	}
}

func main() {
	for _, x := range paths {
		root := ft.NewFileNode(x, x, nil, true)
		fmt.Println(root.AbsPath)
		if tree {
			ft.PrintFileNodeTree(root, []string{}, 1, level, human, md5, imo)
		} else {
			ft.PrintFileNodeSimple(root, human)
		}
		fmt.Println()
	}
}
